// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import (
	"encoding/json"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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

func Test_SqlDatabase_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabase to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabase, SqlDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabase tests if a specific instance of SqlDatabase round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabase(subject SqlDatabase) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabase
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabase
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabase_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabase to SqlDatabase via AssignPropertiesToSqlDatabase & AssignPropertiesFromSqlDatabase returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabase, SqlDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabase tests if a specific instance of SqlDatabase can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabase(subject SqlDatabase) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabase
	err := copied.AssignPropertiesToSqlDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabase
	err = actual.AssignPropertiesFromSqlDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabase, SqlDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabase runs a test to see if a specific instance of SqlDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabase(subject SqlDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabase
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

// Generator of SqlDatabase instances for property testing - lazily instantiated by SqlDatabaseGenerator()
var sqlDatabaseGenerator gopter.Gen

// SqlDatabaseGenerator returns a generator of SqlDatabase instances for property testing.
func SqlDatabaseGenerator() gopter.Gen {
	if sqlDatabaseGenerator != nil {
		return sqlDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabase(generators)
	sqlDatabaseGenerator = gen.Struct(reflect.TypeOf(SqlDatabase{}), generators)

	return sqlDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesSpecGenerator()
	gens["Status"] = SqlDatabaseGetResultsSTATUSGenerator()
}

func Test_DatabaseAccountsSqlDatabases_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabases_Spec to DatabaseAccountsSqlDatabases_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesSpec & AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesSpec, DatabaseAccountsSqlDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesSpec tests if a specific instance of DatabaseAccountsSqlDatabases_Spec can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesSpec(subject DatabaseAccountsSqlDatabases_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccountsSqlDatabases_Spec
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabases_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccountsSqlDatabases_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabases_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesSpec, DatabaseAccountsSqlDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabases_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesSpec(subject DatabaseAccountsSqlDatabases_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabases_Spec
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

// Generator of DatabaseAccountsSqlDatabases_Spec instances for property testing - lazily instantiated by
// DatabaseAccountsSqlDatabasesSpecGenerator()
var databaseAccountsSqlDatabasesSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesSpecGenerator returns a generator of DatabaseAccountsSqlDatabases_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesSpecGenerator != nil {
		return databaseAccountsSqlDatabasesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(generators)
	databaseAccountsSqlDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabases_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(generators)
	databaseAccountsSqlDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabases_Spec{}), generators)

	return databaseAccountsSqlDatabasesSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseResourceGenerator())
}

func Test_SqlDatabaseGetResults_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseGetResults_STATUS to SqlDatabaseGetResults_STATUS via AssignPropertiesToSqlDatabaseGetResultsSTATUS & AssignPropertiesFromSqlDatabaseGetResultsSTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseGetResultsSTATUS, SqlDatabaseGetResultsSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseGetResultsSTATUS tests if a specific instance of SqlDatabaseGetResults_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseGetResultsSTATUS(subject SqlDatabaseGetResults_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseGetResults_STATUS
	err := copied.AssignPropertiesToSqlDatabaseGetResultsSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseGetResults_STATUS
	err = actual.AssignPropertiesFromSqlDatabaseGetResultsSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseGetResults_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseGetResults_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseGetResultsSTATUS, SqlDatabaseGetResultsSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseGetResultsSTATUS runs a test to see if a specific instance of SqlDatabaseGetResults_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseGetResultsSTATUS(subject SqlDatabaseGetResults_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseGetResults_STATUS
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

// Generator of SqlDatabaseGetResults_STATUS instances for property testing - lazily instantiated by
// SqlDatabaseGetResultsSTATUSGenerator()
var sqlDatabaseGetResultsSTATUSGenerator gopter.Gen

// SqlDatabaseGetResultsSTATUSGenerator returns a generator of SqlDatabaseGetResults_STATUS instances for property testing.
// We first initialize sqlDatabaseGetResultsSTATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseGetResultsSTATUSGenerator() gopter.Gen {
	if sqlDatabaseGetResultsSTATUSGenerator != nil {
		return sqlDatabaseGetResultsSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsSTATUS(generators)
	sqlDatabaseGetResultsSTATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetResults_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsSTATUS(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseGetResultsSTATUS(generators)
	sqlDatabaseGetResultsSTATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetResults_STATUS{}), generators)

	return sqlDatabaseGetResultsSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsSTATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseGetResultsSTATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseGetResultsSTATUS(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResourceSTATUSGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseGetPropertiesSTATUSResourceGenerator())
}

func Test_SqlDatabaseGetProperties_STATUS_Resource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseGetProperties_STATUS_Resource to SqlDatabaseGetProperties_STATUS_Resource via AssignPropertiesToSqlDatabaseGetPropertiesSTATUSResource & AssignPropertiesFromSqlDatabaseGetPropertiesSTATUSResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseGetPropertiesSTATUSResource, SqlDatabaseGetPropertiesSTATUSResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseGetPropertiesSTATUSResource tests if a specific instance of SqlDatabaseGetProperties_STATUS_Resource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseGetPropertiesSTATUSResource(subject SqlDatabaseGetProperties_STATUS_Resource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseGetProperties_STATUS_Resource
	err := copied.AssignPropertiesToSqlDatabaseGetPropertiesSTATUSResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseGetProperties_STATUS_Resource
	err = actual.AssignPropertiesFromSqlDatabaseGetPropertiesSTATUSResource(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseGetProperties_STATUS_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseGetProperties_STATUS_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseGetPropertiesSTATUSResource, SqlDatabaseGetPropertiesSTATUSResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseGetPropertiesSTATUSResource runs a test to see if a specific instance of SqlDatabaseGetProperties_STATUS_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseGetPropertiesSTATUSResource(subject SqlDatabaseGetProperties_STATUS_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseGetProperties_STATUS_Resource
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

// Generator of SqlDatabaseGetProperties_STATUS_Resource instances for property testing - lazily instantiated by
// SqlDatabaseGetPropertiesSTATUSResourceGenerator()
var sqlDatabaseGetPropertiesSTATUSResourceGenerator gopter.Gen

// SqlDatabaseGetPropertiesSTATUSResourceGenerator returns a generator of SqlDatabaseGetProperties_STATUS_Resource instances for property testing.
func SqlDatabaseGetPropertiesSTATUSResourceGenerator() gopter.Gen {
	if sqlDatabaseGetPropertiesSTATUSResourceGenerator != nil {
		return sqlDatabaseGetPropertiesSTATUSResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetPropertiesSTATUSResource(generators)
	sqlDatabaseGetPropertiesSTATUSResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetProperties_STATUS_Resource{}), generators)

	return sqlDatabaseGetPropertiesSTATUSResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseGetPropertiesSTATUSResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseGetPropertiesSTATUSResource(gens map[string]gopter.Gen) {
	gens["Colls"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
	gens["Users"] = gen.PtrOf(gen.AlphaString())
}

func Test_SqlDatabaseResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseResource to SqlDatabaseResource via AssignPropertiesToSqlDatabaseResource & AssignPropertiesFromSqlDatabaseResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseResource, SqlDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseResource tests if a specific instance of SqlDatabaseResource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseResource(subject SqlDatabaseResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseResource
	err := copied.AssignPropertiesToSqlDatabaseResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseResource
	err = actual.AssignPropertiesFromSqlDatabaseResource(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseResource, SqlDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseResource runs a test to see if a specific instance of SqlDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseResource(subject SqlDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseResource
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

// Generator of SqlDatabaseResource instances for property testing - lazily instantiated by
// SqlDatabaseResourceGenerator()
var sqlDatabaseResourceGenerator gopter.Gen

// SqlDatabaseResourceGenerator returns a generator of SqlDatabaseResource instances for property testing.
func SqlDatabaseResourceGenerator() gopter.Gen {
	if sqlDatabaseResourceGenerator != nil {
		return sqlDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseResource(generators)
	sqlDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseResource{}), generators)

	return sqlDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}