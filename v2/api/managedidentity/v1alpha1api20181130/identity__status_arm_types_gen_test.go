// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

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

func Test_Identity_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityStatusARM, IdentityStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityStatusARM runs a test to see if a specific instance of Identity_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityStatusARM(subject Identity_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_StatusARM
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

// Generator of Identity_StatusARM instances for property testing - lazily instantiated by IdentityStatusARMGenerator()
var identityStatusARMGenerator gopter.Gen

// IdentityStatusARMGenerator returns a generator of Identity_StatusARM instances for property testing.
// We first initialize identityStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IdentityStatusARMGenerator() gopter.Gen {
	if identityStatusARMGenerator != nil {
		return identityStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityStatusARM(generators)
	identityStatusARMGenerator = gen.Struct(reflect.TypeOf(Identity_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityStatusARM(generators)
	AddRelatedPropertyGeneratorsForIdentityStatusARM(generators)
	identityStatusARMGenerator = gen.Struct(reflect.TypeOf(Identity_StatusARM{}), generators)

	return identityStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentityStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIdentityStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentityStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(UserAssignedIdentityPropertiesStatusARMGenerator())
}

func Test_UserAssignedIdentityProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityPropertiesStatusARM, UserAssignedIdentityPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityPropertiesStatusARM runs a test to see if a specific instance of UserAssignedIdentityProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityPropertiesStatusARM(subject UserAssignedIdentityProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_StatusARM
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

// Generator of UserAssignedIdentityProperties_StatusARM instances for property testing - lazily instantiated by
// UserAssignedIdentityPropertiesStatusARMGenerator()
var userAssignedIdentityPropertiesStatusARMGenerator gopter.Gen

// UserAssignedIdentityPropertiesStatusARMGenerator returns a generator of UserAssignedIdentityProperties_StatusARM instances for property testing.
func UserAssignedIdentityPropertiesStatusARMGenerator() gopter.Gen {
	if userAssignedIdentityPropertiesStatusARMGenerator != nil {
		return userAssignedIdentityPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesStatusARM(generators)
	userAssignedIdentityPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_StatusARM{}), generators)

	return userAssignedIdentityPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}