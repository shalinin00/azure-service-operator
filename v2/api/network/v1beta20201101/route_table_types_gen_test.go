// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import (
	"encoding/json"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
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

func Test_RouteTable_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTable to hub returns original",
		prop.ForAll(RunResourceConversionTestForRouteTable, RouteTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRouteTable tests if a specific instance of RouteTable round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRouteTable(subject RouteTable) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.RouteTable
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RouteTable
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

func Test_RouteTable_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTable to RouteTable via AssignPropertiesToRouteTable & AssignPropertiesFromRouteTable returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTable, RouteTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTable tests if a specific instance of RouteTable can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForRouteTable(subject RouteTable) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTable
	err := copied.AssignPropertiesToRouteTable(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTable
	err = actual.AssignPropertiesFromRouteTable(&other)
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

func Test_RouteTable_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable, RouteTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable runs a test to see if a specific instance of RouteTable round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable(subject RouteTable) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable
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

// Generator of RouteTable instances for property testing - lazily instantiated by RouteTableGenerator()
var routeTableGenerator gopter.Gen

// RouteTableGenerator returns a generator of RouteTable instances for property testing.
func RouteTableGenerator() gopter.Gen {
	if routeTableGenerator != nil {
		return routeTableGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRouteTable(generators)
	routeTableGenerator = gen.Struct(reflect.TypeOf(RouteTable{}), generators)

	return routeTableGenerator
}

// AddRelatedPropertyGeneratorsForRouteTable is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable(gens map[string]gopter.Gen) {
	gens["Spec"] = RouteTablesSpecGenerator()
	gens["Status"] = RouteTableSTATUSGenerator()
}

func Test_RouteTable_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTable_STATUS to RouteTable_STATUS via AssignPropertiesToRouteTableSTATUS & AssignPropertiesFromRouteTableSTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTableSTATUS, RouteTableSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTableSTATUS tests if a specific instance of RouteTable_STATUS can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForRouteTableSTATUS(subject RouteTable_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTable_STATUS
	err := copied.AssignPropertiesToRouteTableSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTable_STATUS
	err = actual.AssignPropertiesFromRouteTableSTATUS(&other)
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

func Test_RouteTable_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTableSTATUS, RouteTableSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTableSTATUS runs a test to see if a specific instance of RouteTable_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTableSTATUS(subject RouteTable_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_STATUS
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

// Generator of RouteTable_STATUS instances for property testing - lazily instantiated by RouteTableSTATUSGenerator()
var routeTableSTATUSGenerator gopter.Gen

// RouteTableSTATUSGenerator returns a generator of RouteTable_STATUS instances for property testing.
func RouteTableSTATUSGenerator() gopter.Gen {
	if routeTableSTATUSGenerator != nil {
		return routeTableSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTableSTATUS(generators)
	routeTableSTATUSGenerator = gen.Struct(reflect.TypeOf(RouteTable_STATUS{}), generators)

	return routeTableSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForRouteTableSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTableSTATUS(gens map[string]gopter.Gen) {
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_RouteTables_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTables_Spec to RouteTables_Spec via AssignPropertiesToRouteTablesSpec & AssignPropertiesFromRouteTablesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTablesSpec, RouteTablesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTablesSpec tests if a specific instance of RouteTables_Spec can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForRouteTablesSpec(subject RouteTables_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTables_Spec
	err := copied.AssignPropertiesToRouteTablesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTables_Spec
	err = actual.AssignPropertiesFromRouteTablesSpec(&other)
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

func Test_RouteTables_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTables_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesSpec, RouteTablesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesSpec runs a test to see if a specific instance of RouteTables_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesSpec(subject RouteTables_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTables_Spec
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

// Generator of RouteTables_Spec instances for property testing - lazily instantiated by RouteTablesSpecGenerator()
var routeTablesSpecGenerator gopter.Gen

// RouteTablesSpecGenerator returns a generator of RouteTables_Spec instances for property testing.
func RouteTablesSpecGenerator() gopter.Gen {
	if routeTablesSpecGenerator != nil {
		return routeTablesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablesSpec(generators)
	routeTablesSpecGenerator = gen.Struct(reflect.TypeOf(RouteTables_Spec{}), generators)

	return routeTablesSpecGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}