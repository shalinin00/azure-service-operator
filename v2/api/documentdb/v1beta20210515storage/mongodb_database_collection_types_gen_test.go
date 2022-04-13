// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_MongodbDatabaseCollection_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollection via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollection, MongodbDatabaseCollectionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollection runs a test to see if a specific instance of MongodbDatabaseCollection round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollection(subject MongodbDatabaseCollection) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollection
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongodbDatabaseCollection instances for property testing - lazily instantiated by
//MongodbDatabaseCollectionGenerator()
var mongodbDatabaseCollectionGenerator gopter.Gen

// MongodbDatabaseCollectionGenerator returns a generator of MongodbDatabaseCollection instances for property testing.
func MongodbDatabaseCollectionGenerator() gopter.Gen {
	if mongodbDatabaseCollectionGenerator != nil {
		return mongodbDatabaseCollectionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollection(generators)
	mongodbDatabaseCollectionGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollection{}), generators)

	return mongodbDatabaseCollectionGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollection is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollection(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsMongodbDatabasesCollectionsSpecGenerator()
	gens["Status"] = MongoDBCollectionGetResultsStatusGenerator()
}

func Test_DatabaseAccountsMongodbDatabasesCollections_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollections_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSpec, DatabaseAccountsMongodbDatabasesCollectionsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSpec runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollections_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollectionsSpec(subject DatabaseAccountsMongodbDatabasesCollections_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollections_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DatabaseAccountsMongodbDatabasesCollections_Spec instances for property testing - lazily instantiated by
//DatabaseAccountsMongodbDatabasesCollectionsSpecGenerator()
var databaseAccountsMongodbDatabasesCollectionsSpecGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollectionsSpecGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollections_Spec instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollectionsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollectionsSpecGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollectionsSpecGenerator != nil {
		return databaseAccountsMongodbDatabasesCollectionsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpec(generators)
	databaseAccountsMongodbDatabasesCollectionsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpec(generators)
	databaseAccountsMongodbDatabasesCollectionsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_Spec{}), generators)

	return databaseAccountsMongodbDatabasesCollectionsSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollectionsSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionResourceGenerator())
}

func Test_MongoDBCollectionGetResults_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetResults_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetResultsStatus, MongoDBCollectionGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetResultsStatus runs a test to see if a specific instance of MongoDBCollectionGetResults_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetResultsStatus(subject MongoDBCollectionGetResults_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetResults_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoDBCollectionGetResults_Status instances for property testing - lazily instantiated by
//MongoDBCollectionGetResultsStatusGenerator()
var mongoDBCollectionGetResultsStatusGenerator gopter.Gen

// MongoDBCollectionGetResultsStatusGenerator returns a generator of MongoDBCollectionGetResults_Status instances for property testing.
// We first initialize mongoDBCollectionGetResultsStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionGetResultsStatusGenerator() gopter.Gen {
	if mongoDBCollectionGetResultsStatusGenerator != nil {
		return mongoDBCollectionGetResultsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatus(generators)
	mongoDBCollectionGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetResults_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatus(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsStatus(generators)
	mongoDBCollectionGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetResults_Status{}), generators)

	return mongoDBCollectionGetResultsStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResourceStatusGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionGetPropertiesStatusResourceGenerator())
}

func Test_MongoDBCollectionGetProperties_Status_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetProperties_Status_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusResource, MongoDBCollectionGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusResource runs a test to see if a specific instance of MongoDBCollectionGetProperties_Status_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetPropertiesStatusResource(subject MongoDBCollectionGetProperties_Status_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetProperties_Status_Resource
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoDBCollectionGetProperties_Status_Resource instances for property testing - lazily instantiated by
//MongoDBCollectionGetPropertiesStatusResourceGenerator()
var mongoDBCollectionGetPropertiesStatusResourceGenerator gopter.Gen

// MongoDBCollectionGetPropertiesStatusResourceGenerator returns a generator of MongoDBCollectionGetProperties_Status_Resource instances for property testing.
// We first initialize mongoDBCollectionGetPropertiesStatusResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionGetPropertiesStatusResourceGenerator() gopter.Gen {
	if mongoDBCollectionGetPropertiesStatusResourceGenerator != nil {
		return mongoDBCollectionGetPropertiesStatusResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResource(generators)
	mongoDBCollectionGetPropertiesStatusResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_Status_Resource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResource(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResource(generators)
	mongoDBCollectionGetPropertiesStatusResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_Status_Resource{}), generators)

	return mongoDBCollectionGetPropertiesStatusResourceGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResource(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetPropertiesStatusResource(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexStatusGenerator())
}

func Test_MongoDBCollectionResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResource, MongoDBCollectionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResource runs a test to see if a specific instance of MongoDBCollectionResource round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResource(subject MongoDBCollectionResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResource
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoDBCollectionResource instances for property testing - lazily instantiated by
//MongoDBCollectionResourceGenerator()
var mongoDBCollectionResourceGenerator gopter.Gen

// MongoDBCollectionResourceGenerator returns a generator of MongoDBCollectionResource instances for property testing.
// We first initialize mongoDBCollectionResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResourceGenerator() gopter.Gen {
	if mongoDBCollectionResourceGenerator != nil {
		return mongoDBCollectionResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource(generators)
	mongoDBCollectionResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResource(generators)
	mongoDBCollectionResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource{}), generators)

	return mongoDBCollectionResourceGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResource(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResource(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexGenerator())
}

func Test_MongoIndex_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndex, MongoIndexGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndex runs a test to see if a specific instance of MongoIndex round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndex(subject MongoIndex) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoIndex instances for property testing - lazily instantiated by MongoIndexGenerator()
var mongoIndexGenerator gopter.Gen

// MongoIndexGenerator returns a generator of MongoIndex instances for property testing.
func MongoIndexGenerator() gopter.Gen {
	if mongoIndexGenerator != nil {
		return mongoIndexGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndex(generators)
	mongoIndexGenerator = gen.Struct(reflect.TypeOf(MongoIndex{}), generators)

	return mongoIndexGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndex is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndex(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsGenerator())
}

func Test_MongoIndex_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexStatus, MongoIndexStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexStatus runs a test to see if a specific instance of MongoIndex_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexStatus(subject MongoIndex_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoIndex_Status instances for property testing - lazily instantiated by MongoIndexStatusGenerator()
var mongoIndexStatusGenerator gopter.Gen

// MongoIndexStatusGenerator returns a generator of MongoIndex_Status instances for property testing.
func MongoIndexStatusGenerator() gopter.Gen {
	if mongoIndexStatusGenerator != nil {
		return mongoIndexStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndexStatus(generators)
	mongoIndexStatusGenerator = gen.Struct(reflect.TypeOf(MongoIndex_Status{}), generators)

	return mongoIndexStatusGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndexStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndexStatus(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysStatusGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsStatusGenerator())
}

func Test_MongoIndexKeys_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeys, MongoIndexKeysGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeys runs a test to see if a specific instance of MongoIndexKeys round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeys(subject MongoIndexKeys) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoIndexKeys instances for property testing - lazily instantiated by MongoIndexKeysGenerator()
var mongoIndexKeysGenerator gopter.Gen

// MongoIndexKeysGenerator returns a generator of MongoIndexKeys instances for property testing.
func MongoIndexKeysGenerator() gopter.Gen {
	if mongoIndexKeysGenerator != nil {
		return mongoIndexKeysGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeys(generators)
	mongoIndexKeysGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys{}), generators)

	return mongoIndexKeysGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeys is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeys(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexKeys_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeysStatus, MongoIndexKeysStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeysStatus runs a test to see if a specific instance of MongoIndexKeys_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeysStatus(subject MongoIndexKeys_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoIndexKeys_Status instances for property testing - lazily instantiated by
//MongoIndexKeysStatusGenerator()
var mongoIndexKeysStatusGenerator gopter.Gen

// MongoIndexKeysStatusGenerator returns a generator of MongoIndexKeys_Status instances for property testing.
func MongoIndexKeysStatusGenerator() gopter.Gen {
	if mongoIndexKeysStatusGenerator != nil {
		return mongoIndexKeysStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeysStatus(generators)
	mongoIndexKeysStatusGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_Status{}), generators)

	return mongoIndexKeysStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeysStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeysStatus(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptions, MongoIndexOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptions runs a test to see if a specific instance of MongoIndexOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptions(subject MongoIndexOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoIndexOptions instances for property testing - lazily instantiated by MongoIndexOptionsGenerator()
var mongoIndexOptionsGenerator gopter.Gen

// MongoIndexOptionsGenerator returns a generator of MongoIndexOptions instances for property testing.
func MongoIndexOptionsGenerator() gopter.Gen {
	if mongoIndexOptionsGenerator != nil {
		return mongoIndexOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptions(generators)
	mongoIndexOptionsGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions{}), generators)

	return mongoIndexOptionsGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptions(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}

func Test_MongoIndexOptions_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptionsStatus, MongoIndexOptionsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptionsStatus runs a test to see if a specific instance of MongoIndexOptions_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptionsStatus(subject MongoIndexOptions_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of MongoIndexOptions_Status instances for property testing - lazily instantiated by
//MongoIndexOptionsStatusGenerator()
var mongoIndexOptionsStatusGenerator gopter.Gen

// MongoIndexOptionsStatusGenerator returns a generator of MongoIndexOptions_Status instances for property testing.
func MongoIndexOptionsStatusGenerator() gopter.Gen {
	if mongoIndexOptionsStatusGenerator != nil {
		return mongoIndexOptionsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptionsStatus(generators)
	mongoIndexOptionsStatusGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_Status{}), generators)

	return mongoIndexOptionsStatusGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptionsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptionsStatus(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}