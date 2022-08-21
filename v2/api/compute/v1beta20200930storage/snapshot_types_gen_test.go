// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200930storage

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

func Test_Snapshot_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot runs a test to see if a specific instance of Snapshot round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot(subject Snapshot) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot
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

// Generator of Snapshot instances for property testing - lazily instantiated by SnapshotGenerator()
var snapshotGenerator gopter.Gen

// SnapshotGenerator returns a generator of Snapshot instances for property testing.
func SnapshotGenerator() gopter.Gen {
	if snapshotGenerator != nil {
		return snapshotGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSnapshot(generators)
	snapshotGenerator = gen.Struct(reflect.TypeOf(Snapshot{}), generators)

	return snapshotGenerator
}

// AddRelatedPropertyGeneratorsForSnapshot is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot(gens map[string]gopter.Gen) {
	gens["Spec"] = SnapshotsSpecGenerator()
	gens["Status"] = SnapshotSTATUSGenerator()
}

func Test_Snapshot_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSTATUS, SnapshotSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSTATUS runs a test to see if a specific instance of Snapshot_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSTATUS(subject Snapshot_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_STATUS
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

// Generator of Snapshot_STATUS instances for property testing - lazily instantiated by SnapshotSTATUSGenerator()
var snapshotSTATUSGenerator gopter.Gen

// SnapshotSTATUSGenerator returns a generator of Snapshot_STATUS instances for property testing.
// We first initialize snapshotSTATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotSTATUSGenerator() gopter.Gen {
	if snapshotSTATUSGenerator != nil {
		return snapshotSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSTATUS(generators)
	snapshotSTATUSGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSTATUS(generators)
	AddRelatedPropertyGeneratorsForSnapshotSTATUS(generators)
	snapshotSTATUSGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUS{}), generators)

	return snapshotSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSTATUS(gens map[string]gopter.Gen) {
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeBytes"] = gen.PtrOf(gen.Int())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.AlphaString())
	gens["HyperVGeneration"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UniqueId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotSTATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotSTATUS(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationDataSTATUSGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionSTATUSGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionSTATUSGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationSTATUSGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanSTATUSGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuSTATUSGenerator())
}

func Test_Snapshots_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshots_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotsSpec, SnapshotsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotsSpec runs a test to see if a specific instance of Snapshots_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotsSpec(subject Snapshots_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshots_Spec
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

// Generator of Snapshots_Spec instances for property testing - lazily instantiated by SnapshotsSpecGenerator()
var snapshotsSpecGenerator gopter.Gen

// SnapshotsSpecGenerator returns a generator of Snapshots_Spec instances for property testing.
// We first initialize snapshotsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SnapshotsSpecGenerator() gopter.Gen {
	if snapshotsSpecGenerator != nil {
		return snapshotsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotsSpec(generators)
	snapshotsSpecGenerator = gen.Struct(reflect.TypeOf(Snapshots_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotsSpec(generators)
	AddRelatedPropertyGeneratorsForSnapshotsSpec(generators)
	snapshotsSpecGenerator = gen.Struct(reflect.TypeOf(Snapshots_Spec{}), generators)

	return snapshotsSpecGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.AlphaString())
	gens["HyperVGeneration"] = gen.PtrOf(gen.AlphaString())
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshotsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshotsSpec(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationDataGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuGenerator())
}

func Test_SnapshotSku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku, SnapshotSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku runs a test to see if a specific instance of SnapshotSku round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku(subject SnapshotSku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku
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

// Generator of SnapshotSku instances for property testing - lazily instantiated by SnapshotSkuGenerator()
var snapshotSkuGenerator gopter.Gen

// SnapshotSkuGenerator returns a generator of SnapshotSku instances for property testing.
func SnapshotSkuGenerator() gopter.Gen {
	if snapshotSkuGenerator != nil {
		return snapshotSkuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku(generators)
	snapshotSkuGenerator = gen.Struct(reflect.TypeOf(SnapshotSku{}), generators)

	return snapshotSkuGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_SnapshotSku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSkuSTATUS, SnapshotSkuSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSkuSTATUS runs a test to see if a specific instance of SnapshotSku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSkuSTATUS(subject SnapshotSku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku_STATUS
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

// Generator of SnapshotSku_STATUS instances for property testing - lazily instantiated by SnapshotSkuSTATUSGenerator()
var snapshotSkuSTATUSGenerator gopter.Gen

// SnapshotSkuSTATUSGenerator returns a generator of SnapshotSku_STATUS instances for property testing.
func SnapshotSkuSTATUSGenerator() gopter.Gen {
	if snapshotSkuSTATUSGenerator != nil {
		return snapshotSkuSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSkuSTATUS(generators)
	snapshotSkuSTATUSGenerator = gen.Struct(reflect.TypeOf(SnapshotSku_STATUS{}), generators)

	return snapshotSkuSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSkuSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSkuSTATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}