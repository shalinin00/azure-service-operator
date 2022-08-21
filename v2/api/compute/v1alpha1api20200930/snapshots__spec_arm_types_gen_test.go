// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

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

func Test_Snapshots_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshots_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotsSpecARM, SnapshotsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotsSpecARM runs a test to see if a specific instance of Snapshots_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotsSpecARM(subject Snapshots_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshots_SpecARM
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

// Generator of Snapshots_SpecARM instances for property testing - lazily instantiated by SnapshotsSpecARMGenerator()
var snapshotsSpecARMGenerator gopter.Gen

// SnapshotsSpecARMGenerator returns a generator of Snapshots_SpecARM instances for property testing.
// We first initialize snapshotsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotsSpecARMGenerator() gopter.Gen {
	if snapshotsSpecARMGenerator != nil {
		return snapshotsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotsSpecARM(generators)
	snapshotsSpecARMGenerator = gen.Struct(reflect.TypeOf(Snapshots_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotsSpecARM(generators)
	AddRelatedPropertyGeneratorsForSnapshotsSpecARM(generators)
	snapshotsSpecARMGenerator = gen.Struct(reflect.TypeOf(Snapshots_SpecARM{}), generators)

	return snapshotsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotsSpecARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationARMGenerator())
	gens["Properties"] = gen.PtrOf(SnapshotPropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuARMGenerator())
}

func Test_SnapshotPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotPropertiesARM, SnapshotPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotPropertiesARM runs a test to see if a specific instance of SnapshotPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotPropertiesARM(subject SnapshotPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotPropertiesARM
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

// Generator of SnapshotPropertiesARM instances for property testing - lazily instantiated by
// SnapshotPropertiesARMGenerator()
var snapshotPropertiesARMGenerator gopter.Gen

// SnapshotPropertiesARMGenerator returns a generator of SnapshotPropertiesARM instances for property testing.
// We first initialize snapshotPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotPropertiesARMGenerator() gopter.Gen {
	if snapshotPropertiesARMGenerator != nil {
		return snapshotPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotPropertiesARM(generators)
	snapshotPropertiesARMGenerator = gen.Struct(reflect.TypeOf(SnapshotPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForSnapshotPropertiesARM(generators)
	snapshotPropertiesARMGenerator = gen.Struct(reflect.TypeOf(SnapshotPropertiesARM{}), generators)

	return snapshotPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotPropertiesARM(gens map[string]gopter.Gen) {
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		SnapshotPropertiesDiskState_ActiveSAS,
		SnapshotPropertiesDiskState_ActiveUpload,
		SnapshotPropertiesDiskState_Attached,
		SnapshotPropertiesDiskState_ReadyToUpload,
		SnapshotPropertiesDiskState_Reserved,
		SnapshotPropertiesDiskState_Unattached))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesHyperVGeneration_V1, SnapshotPropertiesHyperVGeneration_V2))
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesNetworkAccessPolicy_AllowAll, SnapshotPropertiesNetworkAccessPolicy_AllowPrivate, SnapshotPropertiesNetworkAccessPolicy_DenyAll))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotPropertiesOsType_Linux, SnapshotPropertiesOsType_Windows))
}

// AddRelatedPropertyGeneratorsForSnapshotPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotPropertiesARM(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationDataARMGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionARMGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionARMGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanARMGenerator())
}

func Test_SnapshotSkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSkuARM, SnapshotSkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSkuARM runs a test to see if a specific instance of SnapshotSkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSkuARM(subject SnapshotSkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSkuARM
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

// Generator of SnapshotSkuARM instances for property testing - lazily instantiated by SnapshotSkuARMGenerator()
var snapshotSkuARMGenerator gopter.Gen

// SnapshotSkuARMGenerator returns a generator of SnapshotSkuARM instances for property testing.
func SnapshotSkuARMGenerator() gopter.Gen {
	if snapshotSkuARMGenerator != nil {
		return snapshotSkuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSkuARM(generators)
	snapshotSkuARMGenerator = gen.Struct(reflect.TypeOf(SnapshotSkuARM{}), generators)

	return snapshotSkuARMGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSkuARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSkuName_PremiumLRS, SnapshotSkuName_StandardLRS, SnapshotSkuName_StandardZRS))
}