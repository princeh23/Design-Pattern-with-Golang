package main

import (
	"fmt"
	"path"
)

type IResourceFile interface {
	Accept(Visitor) error
}

type PdfFile struct {
	path string
}

func (f *PdfFile) Accept(visitor Visitor) error {
	return visitor.Visit(f)
}

type PPTFile struct {
	path string
}

func (f *PPTFile) Accept(visitor Visitor) error {
	return visitor.Visit(f)
}

func NewResourceFile(filepath string) (IResourceFile, error) {
	switch path.Ext(filepath) {
	case ".ppt":
		return &PPTFile{path: filepath}, nil
	case ".pdf":
		return &PdfFile{path: filepath}, nil
	default:
		return nil, fmt.Errorf("not found file type: %s", filepath)
	}
}

type Visitor interface {
	Visit(IResourceFile) error
}

type Compressor struct{}

func (c *Compressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PPTFile:
		return c.CompressPPTFile(f)
	case *PdfFile:
		return c.CompressPDFFile(f)
	default:
		return fmt.Errorf("not found resource typr: %#v", r)
	}
}

func (c *Compressor) CompressPPTFile(f *PPTFile) error {
	fmt.Println("compress ppt file completely")
	return nil
}

func (c *Compressor) CompressPDFFile(f *PdfFile) error {
	fmt.Println("compress pdf file completely")
	return nil
}

type Decompressor struct{}

func (d *Decompressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PPTFile:
		return d.DecompressPPTFile(f)
	case *PdfFile:
		return d.DecompressPDFFile(f)
	default:
		return fmt.Errorf("not found resource typr: %#v", r)
	}
}

func (d *Decompressor) DecompressPPTFile(f *PPTFile) error {
	fmt.Println("decompress ppt file completely")
	return nil
}

func (d *Decompressor) DecompressPDFFile(f *PdfFile) error {
	fmt.Println("decompress pdf file completely")
	return nil
}

func main() {
	pptFile, _ := NewResourceFile(".ppt")
	pdfFile, _ := NewResourceFile(".pdf")

	var visitor Visitor
	visitor = &Compressor{}

	pptFile.Accept(visitor)
	pdfFile.Accept(visitor)

	visitor = &Decompressor{}

	pptFile.Accept(visitor)
	pdfFile.Accept(visitor)
}
