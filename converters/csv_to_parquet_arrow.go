package converters

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet"
	"github.com/apache/arrow-go/v18/parquet/compress"
	pqarrow "github.com/apache/arrow-go/v18/parquet/pqarrow"
)

func inferArrowType(val string) arrow.DataType {
	if _, err := strconv.ParseInt(val, 10, 32); err == nil {
		return arrow.PrimitiveTypes.Int32
	}
	if _, err := strconv.ParseFloat(val, 64); err == nil {
		return arrow.PrimitiveTypes.Float64
	}
	if val == "true" || val == "false" {
		return arrow.FixedWidthTypes.Boolean
	}
	return arrow.BinaryTypes.String
}

func CSVtoParquetArrow(inputCSV, outputParquet string) error {
	f, err := os.Open(inputCSV)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	headers, err := reader.Read()
	if err != nil {
		return err
	}

	records, err := reader.ReadAll()
	if err != nil && err != io.EOF {
		return err
	}

	sample := records[0]
	fields := make([]arrow.Field, len(headers))
	for i, name := range headers {
		dt := inferArrowType(sample[i])
		fields[i] = arrow.Field{Name: name, Type: dt, Nullable: false}
	}
	schema := *arrow.NewSchema(fields, nil)

	pool := memory.NewGoAllocator()
	bldr := array.NewRecordBuilder(pool, &schema)
	defer bldr.Release()

	for _, row := range records {
		for i, cell := range row {
			switch schema.Field(i).Type.ID() {
			case arrow.INT32:
				v, _ := strconv.ParseInt(cell, 10, 32)
				bldr.Field(i).(*array.Int32Builder).Append(int32(v))
			case arrow.FLOAT64:
				fv, _ := strconv.ParseFloat(cell, 64)
				bldr.Field(i).(*array.Float64Builder).Append(fv)
			case arrow.BOOL:
				b, _ := strconv.ParseBool(cell)
				bldr.Field(i).(*array.BooleanBuilder).Append(b)
			default:
				bldr.Field(i).(*array.StringBuilder).Append(cell)
			}
		}
	}

	rec := bldr.NewRecord()
	defer rec.Release()

	out, err := os.Create(outputParquet)
	if err != nil {
		return err
	}
	defer out.Close()

	pProps := parquet.NewWriterProperties(
		parquet.WithCompression(compress.Codecs.Snappy),
	)
	aProps := pqarrow.DefaultWriterProps()

	pw, err := pqarrow.NewFileWriter(&schema, out, pProps, aProps)
	if err != nil {
		return err
	}
	defer pw.Close()

	if err := pw.Write(rec); err != nil {
		return err
	}

	return nil
}
