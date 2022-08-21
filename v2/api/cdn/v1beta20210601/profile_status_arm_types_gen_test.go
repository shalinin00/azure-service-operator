// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

func Test_Profile_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profile_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileSTATUSARM, ProfileSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileSTATUSARM runs a test to see if a specific instance of Profile_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileSTATUSARM(subject Profile_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profile_STATUSARM
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

// Generator of Profile_STATUSARM instances for property testing - lazily instantiated by ProfileSTATUSARMGenerator()
var profileSTATUSARMGenerator gopter.Gen

// ProfileSTATUSARMGenerator returns a generator of Profile_STATUSARM instances for property testing.
// We first initialize profileSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProfileSTATUSARMGenerator() gopter.Gen {
	if profileSTATUSARMGenerator != nil {
		return profileSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileSTATUSARM(generators)
	profileSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Profile_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForProfileSTATUSARM(generators)
	profileSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Profile_STATUSARM{}), generators)

	return profileSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForProfileSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfileSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfileSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfileSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProfilePropertiesSTATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuSTATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataSTATUSARMGenerator())
}

func Test_ProfileProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProfileProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfilePropertiesSTATUSARM, ProfilePropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfilePropertiesSTATUSARM runs a test to see if a specific instance of ProfileProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfilePropertiesSTATUSARM(subject ProfileProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProfileProperties_STATUSARM
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

// Generator of ProfileProperties_STATUSARM instances for property testing - lazily instantiated by
// ProfilePropertiesSTATUSARMGenerator()
var profilePropertiesSTATUSARMGenerator gopter.Gen

// ProfilePropertiesSTATUSARMGenerator returns a generator of ProfileProperties_STATUSARM instances for property testing.
func ProfilePropertiesSTATUSARMGenerator() gopter.Gen {
	if profilePropertiesSTATUSARMGenerator != nil {
		return profilePropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfilePropertiesSTATUSARM(generators)
	profilePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ProfileProperties_STATUSARM{}), generators)

	return profilePropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForProfilePropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfilePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["FrontDoorId"] = gen.PtrOf(gen.AlphaString())
	gens["OriginResponseTimeoutSeconds"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProfilePropertiesSTATUSProvisioningState_Creating,
		ProfilePropertiesSTATUSProvisioningState_Deleting,
		ProfilePropertiesSTATUSProvisioningState_Failed,
		ProfilePropertiesSTATUSProvisioningState_Succeeded,
		ProfilePropertiesSTATUSProvisioningState_Updating))
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ProfilePropertiesSTATUSResourceState_Active,
		ProfilePropertiesSTATUSResourceState_Creating,
		ProfilePropertiesSTATUSResourceState_Deleting,
		ProfilePropertiesSTATUSResourceState_Disabled))
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuSTATUSARM, SkuSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuSTATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuSTATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by SkuSTATUSARMGenerator()
var skuSTATUSARMGenerator gopter.Gen

// SkuSTATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func SkuSTATUSARMGenerator() gopter.Gen {
	if skuSTATUSARMGenerator != nil {
		return skuSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuSTATUSARM(generators)
	skuSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return skuSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuSTATUSARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		SkuSTATUSName_CustomVerizon,
		SkuSTATUSName_PremiumAzureFrontDoor,
		SkuSTATUSName_PremiumVerizon,
		SkuSTATUSName_Standard955BandWidthChinaCdn,
		SkuSTATUSName_StandardAkamai,
		SkuSTATUSName_StandardAvgBandWidthChinaCdn,
		SkuSTATUSName_StandardAzureFrontDoor,
		SkuSTATUSName_StandardChinaCdn,
		SkuSTATUSName_StandardMicrosoft,
		SkuSTATUSName_StandardPlus955BandWidthChinaCdn,
		SkuSTATUSName_StandardPlusAvgBandWidthChinaCdn,
		SkuSTATUSName_StandardPlusChinaCdn,
		SkuSTATUSName_StandardVerizon))
}