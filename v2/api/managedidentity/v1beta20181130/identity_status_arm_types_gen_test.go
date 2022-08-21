// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130

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

func Test_Identity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentitySTATUSARM, IdentitySTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentitySTATUSARM runs a test to see if a specific instance of Identity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentitySTATUSARM(subject Identity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUSARM
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

// Generator of Identity_STATUSARM instances for property testing - lazily instantiated by IdentitySTATUSARMGenerator()
var identitySTATUSARMGenerator gopter.Gen

// IdentitySTATUSARMGenerator returns a generator of Identity_STATUSARM instances for property testing.
// We first initialize identitySTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IdentitySTATUSARMGenerator() gopter.Gen {
	if identitySTATUSARMGenerator != nil {
		return identitySTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentitySTATUSARM(generators)
	identitySTATUSARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentitySTATUSARM(generators)
	AddRelatedPropertyGeneratorsForIdentitySTATUSARM(generators)
	identitySTATUSARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUSARM{}), generators)

	return identitySTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentitySTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentitySTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIdentitySTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentitySTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(UserAssignedIdentityPropertiesSTATUSARMGenerator())
}

func Test_UserAssignedIdentityProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityPropertiesSTATUSARM, UserAssignedIdentityPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityPropertiesSTATUSARM runs a test to see if a specific instance of UserAssignedIdentityProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityPropertiesSTATUSARM(subject UserAssignedIdentityProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityProperties_STATUSARM
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

// Generator of UserAssignedIdentityProperties_STATUSARM instances for property testing - lazily instantiated by
// UserAssignedIdentityPropertiesSTATUSARMGenerator()
var userAssignedIdentityPropertiesSTATUSARMGenerator gopter.Gen

// UserAssignedIdentityPropertiesSTATUSARMGenerator returns a generator of UserAssignedIdentityProperties_STATUSARM instances for property testing.
func UserAssignedIdentityPropertiesSTATUSARMGenerator() gopter.Gen {
	if userAssignedIdentityPropertiesSTATUSARMGenerator != nil {
		return userAssignedIdentityPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesSTATUSARM(generators)
	userAssignedIdentityPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityProperties_STATUSARM{}), generators)

	return userAssignedIdentityPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentityPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}