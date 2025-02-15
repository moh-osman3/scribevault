package diff

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersionNumber(t *testing.T) {
	v := &Version{versionNumber: 1}
	assert.Equal(t, int64(1), v.getVersionNumber())
}

func TestGetDocumentAsString(t *testing.T) {
	d := &Document{name: "test", header: "header", body: "body"}
	expected := "test:header:body"
	assert.Equal(t, expected, d.getDocumentAsString())
}

func TestAddVersion(t *testing.T) {
	d := &Document{}
	v := &Version{versionNumber: 1}
	d.addVersion(v)
	assert.Len(t, d.getVersions(), 1)
	assert.Equal(t, v, d.getVersions()[0])
}

func TestSaveDocument(t *testing.T) {
	db := make(map[string]*Document)
	dv := &DefaultVersioner{documentDb: db}
	doc := &Document{name: "test"}

	err := dv.saveDocument("test", doc)
	assert.NoError(t, err)
	assert.Equal(t, doc, db["test"])
}

func TestSaveDocumentNoOverride(t *testing.T) {
	db := make(map[string]*Document)
	dv := &DefaultVersioner{documentDb: db}
	doc := &Document{name: "test"}

	err := dv.saveDocumentNoOverride("test", doc)
	assert.NoError(t, err)
	assert.Equal(t, doc, db["test"])

	err = dv.saveDocumentNoOverride("test", doc)
	assert.Error(t, err)
}

func TestGetDocument(t *testing.T) {
	db := make(map[string]*Document)
	doc := &Document{name: "test"}
	db["test"] = doc
	dv := &DefaultVersioner{documentDb: db}

	result, err := dv.getDocument("test")
	assert.NoError(t, err)
	assert.Equal(t, doc, result)

	_, err = dv.getDocument("nonexistent")
	assert.Error(t, err)
}

func TestGetLatestVersion(t *testing.T) {
	doc := &Document{name: "test"}
	v1 := &Version{versionNumber: 1, document: doc}
	v2 := &Version{versionNumber: 2, document: doc}
	doc.addVersion(v1)
	doc.addVersion(v2)
	db := map[string]*Document{"test": doc}
	dv := &DefaultVersioner{documentDb: db}

	latest, err := dv.getLatestVersion("test")
	assert.NoError(t, err)
	assert.Equal(t, v2, latest)
}

func TestSaveChanges(t *testing.T) {
	db := make(map[string]*Document)
	dv := &DefaultVersioner{documentDb: db}
	doc := &Document{name: "test"}

	dv.saveDocument("test", doc)

	// Ensure document has at least one initial version
	initialVersion := &Version{versionNumber: 1, document: doc}
	doc.addVersion(initialVersion)

	err := dv.saveChanges("test", doc)
	assert.NoError(t, err)

	versions := doc.getVersions()
	assert.Len(t, versions, 2)
	assert.NotNil(t, versions[1])
	assert.Equal(t, int64(2), versions[1].getVersionNumber())
}
