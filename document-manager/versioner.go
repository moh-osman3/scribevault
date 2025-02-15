package diff

import (
	"fmt"
	"time"
)

type Version struct {
	versionNumber int64
	document      *Document
	lastVersion   *Version
}

func (v *Version) getVersionNumber() int64 {
	return v.versionNumber
}

type Document struct {
	name         string
	createdTime  time.Time
	lastModified time.Time
	body         string
	header       string
	versions     []*Version
}

func (d *Document) getDocumentAsString() string {
	return fmt.Sprintf("%s:%s:%s", d.name, d.header, d.body)
}

func (d *Document) getVersions() []*Version {
	return d.versions
}

func (d *Document) addVersion(v *Version) {
	d.versions = append(d.versions, v)
}

type diffFunction func(s1, s2 string) (string, string)

type Versioner interface {
	diffVersions(fn diffFunction, v1, v2 *Version) (plusDoc Document, minusDoc Document)
	getLatestVersion(documentName string) (*Version, error)
	getDocument(documentName string) (*Document, error)
	saveDocument(documentName string, doc *Document) error
	saveChanges(documentName string, doc *Document) error
}

type DefaultVersioner struct {
	documentDb map[string]*Document
}

func (dv *DefaultVersioner) getDocument(documentName string) (*Document, error) {
	doc, ok := dv.documentDb[documentName]
	if !ok {
		return nil, fmt.Errorf("version.go: could not get document: document does not exist: %v", documentName)
	}
	return doc, nil
}

func (dv *DefaultVersioner) saveDocumentNoOverride(documentName string, doc *Document) error {
	if dv.documentDb == nil {
		return fmt.Errorf("version.go: unable to save document: document database does not exist: %v", documentName)
	}
	_, err := dv.getDocument(documentName)
	if err == nil {
		return fmt.Errorf("version.go: could not save document: document already exists: %v", documentName)
	}

	dv.documentDb[documentName] = doc

	return nil
}

func (dv *DefaultVersioner) saveDocument(documentName string, doc *Document) error {
	if dv.documentDb == nil {
		return fmt.Errorf("version.go: unable to save document: document database does not exist: %v", documentName)
	}
	dv.documentDb[documentName] = doc
	return nil
}

func (dv *DefaultVersioner) saveChanges(documentName string, doc *Document) error {
	// check if document already exists
	err := dv.saveDocumentNoOverride(documentName, doc)
	if err == nil {
		return err
	}

	recentVersion, err := dv.getLatestVersion(documentName)
	if err != nil {
		return err
	}

	newVersion := &Version{
		versionNumber: recentVersion.getVersionNumber() + 1,
		document:      doc,
		lastVersion:   recentVersion,
	}

	doc.addVersion(newVersion)
	err = dv.saveDocument(documentName, doc)

	return err
}

func (dv *DefaultVersioner) getLatestVersion(documentName string) (*Version, error) {
	doc, err := dv.getDocument(documentName)

	if err != nil {
		return nil, fmt.Errorf("version.go: could not get latest version: %v", err)
	}

	versionSlice := doc.getVersions()
	length := len(versionSlice)

	if length == 0 {
		return nil, fmt.Errorf("version.go: could not get latest version: no versions found for document: %s", documentName)
	}

	return versionSlice[length-1], nil
}
