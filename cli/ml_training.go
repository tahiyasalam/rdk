package cli

import (
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/types/known/structpb"
)

// MLTrainingUploadAction retrieves the logs for a specific build step.
func MLTrainingUploadAction(c *cli.Context) error {
	client, err := newViamClient(c)
	if err != nil {
		return err
	}

	metadata := createMetadata(c.Bool(mlTrainingFlagDraft), c.String(mlTrainingFlagType),
		c.String(mlTrainingFlagFramework))
	metadataStruct, err := convertMetadataToStruct(metadata)
	if err != nil {
		return err
	}

	resp, err := client.uploadPackage(c.String(generalFlagOrgID),
		c.String(mlTrainingFlagName),
		c.String(mlTrainingFlagVersion),
		string(PackageTypeMLTraining),
		c.Path(mlTrainingFlagPath),
		metadataStruct,
	)
	if err != nil {
		return err
	}

	printf(c.App.Writer, "Version successfully uploaded! you can view your changes online here: %s", resp.GetId())
	return nil
}

// ModelType refers to the type of the model.
type ModelType string

// ModelType enumeration.
const (
	ModelTypeUnspecified               = ModelType("unspecified")
	ModelTypeSingleLabelClassification = ModelType("single_label_classification")
	ModelTypeMultiLabelClassification  = ModelType("multi_label_classification")
	ModelTypeObjectDetection           = ModelType("object_detection")
)

var modelTypes = []string{
	string(ModelTypeUnspecified), string(ModelTypeSingleLabelClassification),
	string(ModelTypeMultiLabelClassification), string(ModelTypeObjectDetection),
}

// ModelFramework refers to the backend framework of the model.
type ModelFramework string

// ModelFramework enumeration.
const (
	ModelFrameworkUnspecified = ModelFramework("unspecified")
	ModelFrameworkTFLite      = ModelFramework("tflite")
	ModelFrameworkTensorFlow  = ModelFramework("tensorflow")
	ModelFrameworkPyTorch     = ModelFramework("py_torch")
	ModelFrameworkONNX        = ModelFramework("onnx")
)

var modelFrameworks = []string{
	string(PackageTypeUnspecified), string(PackageTypeArchive), string(PackageTypeMLModel),
	string(PackageTypeModule), string(PackageTypeSLAMMap),
}

// MLMetadata struct stores package info for ML training packages.
type MLMetadata struct {
	Draft     bool
	ModelType string
	Framework string
}

func createMetadata(draft bool, modelType, framework string) MLMetadata {
	return MLMetadata{
		Draft:     draft,
		ModelType: findValueOrSetDefault(modelTypes, modelType, string(ModelTypeUnspecified)),
		Framework: findValueOrSetDefault(modelFrameworks, framework, string(ModelFrameworkUnspecified)),
	}
}

func findValueOrSetDefault(arr []string, val, defaultVal string) string {
	for _, str := range arr {
		if str == val {
			return val
		}
	}
	return defaultVal
}

var (
	modelTypeKey      = "model_type"
	modelFrameworkKey = "model_framework"
	draftKey          = "draft"
)

func convertMetadataToStruct(metadata MLMetadata) (*structpb.Struct, error) {
	metadataMap := make(map[string]interface{})
	metadataMap[modelTypeKey] = metadata.ModelType
	metadataMap[modelFrameworkKey] = metadata.Framework
	metadataMap[draftKey] = metadata.Draft
	metadataStruct, err := structpb.NewStruct(metadataMap)
	if err != nil {
		return nil, err
	}
	return metadataStruct, nil
}
