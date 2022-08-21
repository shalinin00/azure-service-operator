// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200801previewstorage

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

func Test_RoleAssignment_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment runs a test to see if a specific instance of RoleAssignment round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment(subject RoleAssignment) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment
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

// Generator of RoleAssignment instances for property testing - lazily instantiated by RoleAssignmentGenerator()
var roleAssignmentGenerator gopter.Gen

// RoleAssignmentGenerator returns a generator of RoleAssignment instances for property testing.
func RoleAssignmentGenerator() gopter.Gen {
	if roleAssignmentGenerator != nil {
		return roleAssignmentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRoleAssignment(generators)
	roleAssignmentGenerator = gen.Struct(reflect.TypeOf(RoleAssignment{}), generators)

	return roleAssignmentGenerator
}

// AddRelatedPropertyGeneratorsForRoleAssignment is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment(gens map[string]gopter.Gen) {
	gens["Spec"] = RoleAssignmentsSpecGenerator()
	gens["Status"] = RoleAssignmentSTATUSGenerator()
}

func Test_RoleAssignment_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentSTATUS, RoleAssignmentSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentSTATUS runs a test to see if a specific instance of RoleAssignment_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentSTATUS(subject RoleAssignment_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_STATUS
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

// Generator of RoleAssignment_STATUS instances for property testing - lazily instantiated by
// RoleAssignmentSTATUSGenerator()
var roleAssignmentSTATUSGenerator gopter.Gen

// RoleAssignmentSTATUSGenerator returns a generator of RoleAssignment_STATUS instances for property testing.
func RoleAssignmentSTATUSGenerator() gopter.Gen {
	if roleAssignmentSTATUSGenerator != nil {
		return roleAssignmentSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentSTATUS(generators)
	roleAssignmentSTATUSGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_STATUS{}), generators)

	return roleAssignmentSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentSTATUS(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.AlphaString())
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}

func Test_RoleAssignments_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignments_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentsSpec, RoleAssignmentsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentsSpec runs a test to see if a specific instance of RoleAssignments_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentsSpec(subject RoleAssignments_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignments_Spec
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

// Generator of RoleAssignments_Spec instances for property testing - lazily instantiated by
// RoleAssignmentsSpecGenerator()
var roleAssignmentsSpecGenerator gopter.Gen

// RoleAssignmentsSpecGenerator returns a generator of RoleAssignments_Spec instances for property testing.
func RoleAssignmentsSpecGenerator() gopter.Gen {
	if roleAssignmentsSpecGenerator != nil {
		return roleAssignmentsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentsSpec(generators)
	roleAssignmentsSpecGenerator = gen.Struct(reflect.TypeOf(RoleAssignments_Spec{}), generators)

	return roleAssignmentsSpecGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}