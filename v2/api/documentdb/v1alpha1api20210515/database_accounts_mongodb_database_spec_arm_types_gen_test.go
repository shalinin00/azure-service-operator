// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccounts_MongodbDatabase_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabase_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_SpecARM, DatabaseAccounts_MongodbDatabase_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_SpecARM runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabase_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_SpecARM(subject DatabaseAccounts_MongodbDatabase_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabase_SpecARM
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

// Generator of DatabaseAccounts_MongodbDatabase_SpecARM instances for property testing - lazily instantiated by
// DatabaseAccounts_MongodbDatabase_SpecARMGenerator()
var databaseAccounts_MongodbDatabase_SpecARMGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabase_SpecARMGenerator returns a generator of DatabaseAccounts_MongodbDatabase_SpecARM instances for property testing.
// We first initialize databaseAccounts_MongodbDatabase_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabase_SpecARMGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabase_SpecARMGenerator != nil {
		return databaseAccounts_MongodbDatabase_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_SpecARM(generators)
	databaseAccounts_MongodbDatabase_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_SpecARM(generators)
	databaseAccounts_MongodbDatabase_SpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_SpecARM{}), generators)

	return databaseAccounts_MongodbDatabase_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBDatabaseCreateUpdatePropertiesARMGenerator())
}

func Test_MongoDBDatabaseCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseCreateUpdatePropertiesARM, MongoDBDatabaseCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseCreateUpdatePropertiesARM runs a test to see if a specific instance of MongoDBDatabaseCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseCreateUpdatePropertiesARM(subject MongoDBDatabaseCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseCreateUpdatePropertiesARM
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

// Generator of MongoDBDatabaseCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
// MongoDBDatabaseCreateUpdatePropertiesARMGenerator()
var mongoDBDatabaseCreateUpdatePropertiesARMGenerator gopter.Gen

// MongoDBDatabaseCreateUpdatePropertiesARMGenerator returns a generator of MongoDBDatabaseCreateUpdatePropertiesARM instances for property testing.
func MongoDBDatabaseCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if mongoDBDatabaseCreateUpdatePropertiesARMGenerator != nil {
		return mongoDBDatabaseCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdatePropertiesARM(generators)
	mongoDBDatabaseCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseCreateUpdatePropertiesARM{}), generators)

	return mongoDBDatabaseCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResourceARMGenerator())
}

func Test_CreateUpdateOptionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptionsARM, CreateUpdateOptionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptionsARM runs a test to see if a specific instance of CreateUpdateOptionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptionsARM(subject CreateUpdateOptionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptionsARM
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

// Generator of CreateUpdateOptionsARM instances for property testing - lazily instantiated by
// CreateUpdateOptionsARMGenerator()
var createUpdateOptionsARMGenerator gopter.Gen

// CreateUpdateOptionsARMGenerator returns a generator of CreateUpdateOptionsARM instances for property testing.
// We first initialize createUpdateOptionsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptionsARMGenerator() gopter.Gen {
	if createUpdateOptionsARMGenerator != nil {
		return createUpdateOptionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM(generators)
	createUpdateOptionsARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptionsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptionsARM(generators)
	createUpdateOptionsARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptionsARM{}), generators)

	return createUpdateOptionsARMGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptionsARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptionsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptionsARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsARMGenerator())
}

func Test_MongoDBDatabaseResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseResourceARM, MongoDBDatabaseResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseResourceARM runs a test to see if a specific instance of MongoDBDatabaseResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseResourceARM(subject MongoDBDatabaseResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseResourceARM
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

// Generator of MongoDBDatabaseResourceARM instances for property testing - lazily instantiated by
// MongoDBDatabaseResourceARMGenerator()
var mongoDBDatabaseResourceARMGenerator gopter.Gen

// MongoDBDatabaseResourceARMGenerator returns a generator of MongoDBDatabaseResourceARM instances for property testing.
func MongoDBDatabaseResourceARMGenerator() gopter.Gen {
	if mongoDBDatabaseResourceARMGenerator != nil {
		return mongoDBDatabaseResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResourceARM(generators)
	mongoDBDatabaseResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResourceARM{}), generators)

	return mongoDBDatabaseResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseResourceARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_AutoscaleSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsARM, AutoscaleSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsARM runs a test to see if a specific instance of AutoscaleSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsARM(subject AutoscaleSettingsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsARM
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

// Generator of AutoscaleSettingsARM instances for property testing - lazily instantiated by
// AutoscaleSettingsARMGenerator()
var autoscaleSettingsARMGenerator gopter.Gen

// AutoscaleSettingsARMGenerator returns a generator of AutoscaleSettingsARM instances for property testing.
func AutoscaleSettingsARMGenerator() gopter.Gen {
	if autoscaleSettingsARMGenerator != nil {
		return autoscaleSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsARM(generators)
	autoscaleSettingsARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsARM{}), generators)

	return autoscaleSettingsARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}