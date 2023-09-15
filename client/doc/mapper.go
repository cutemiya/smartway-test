package doc

import (
	"smartway-test/model"
)

func MapToDocumentServiceModel(docModel Document) model.Document {
	return model.Document{
		Id:     docModel.Id,
		Type:   docModel.Type,
		Number: docModel.Number,
	}
}

func MapToDocumentsServiceModel(docsModel DocumentList) model.DocumentList {
	var docList []model.Document
	for _, document := range docsModel.Documents {
		docList = append(docList, MapToDocumentServiceModel(document))
	}

	return model.DocumentList{
		Documents: docList,
	}
}

// reverse map

func MapToClientDocumentModel(docsList []model.Document) []Document {
	var docs []Document

	for _, document := range docsList {
		docs = append(docs, Document{
			Id:     document.Id,
			Type:   document.Type,
			Number: document.Number,
		})
	}

	return docs
}
