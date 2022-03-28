package imageshrink

import (
	"image"
	"os"
	"reflect"
	"testing"
)

type ShrinkData struct {
	inputName    string
	expectedName string
}

var testData = []ShrinkData{
	{"test_data/input_01.png", "test_data/output_input_01.png"},
	{"test_data/input_02.png", "test_data/output_input_02.png"},
	{"test_data/input_03.png", "test_data/output_input_03.png"},
	{"test_data/input_04.png", "test_data/output_input_04.png"},
	{"test_data/input_05.png", "test_data/output_input_05.png"},
	{"test_data/input_06.png", "test_data/output_input_06.png"},
	{"test_data/input_07.png", "test_data/output_input_07.png"},
	{"test_data/input_08.png", "test_data/output_input_08.png"},
	{"test_data/input_09.png", "test_data/output_input_09.png"},
	{"test_data/input_10.png", "test_data/output_input_10.png"},
	{"test_data/input_11.png", "test_data/output_input_11.png"},
	{"test_data/input_12.png", "test_data/output_input_12.png"},
	{"test_data/input_13.png", "test_data/output_input_13.png"},
	{"test_data/input_14.png", "test_data/output_input_14.png"},
	{"test_data/input_15.png", "test_data/output_input_15.png"},
	{"test_data/input_16.png", "test_data/output_input_16.png"},
	{"test_data/input_17.png", "test_data/output_input_17.png"},
	{"test_data/input_18.png", "test_data/output_input_18.png"},
	{"test_data/input_19.png", "test_data/output_input_19.png"},
	{"test_data/input_20.png", "test_data/output_input_20.png"},
}

var invalidInput = []ShrinkData{
	{"ThisDoesn'tExist", "test_data/temp.png"},
	{"test_data/input_01.png", "dir/doesn't/exist/temp.png"},
	{"test_data/not_an_image.png", "test_data/temp.png"},
	{"test_data/opaque.png", "test_data/temp.png"},
}

func CompareImages(src, dst string) (error, bool) {
	fsrc, err := os.Open(src)
	if err != nil {
		return err, false
	}
	fdst, err := os.Open(dst)
	if err != nil {
		return err, false
	}
	isrc, _, err := image.Decode(fsrc)
	if err != nil {
		return err, false
	}
	idst, _, err := image.Decode(fdst)
	if err != nil {
		return err, false
	}
	return nil, reflect.DeepEqual(isrc, idst)
}

func TestShrinkFile(t *testing.T) {
	for _, datum := range testData {
		err := ShrinkFile(datum.inputName, "test_data/temp.png")
		if err != nil {
			t.Error(err)
		}
		err, result := CompareImages(datum.expectedName, "test_data/temp.png")
		if err != nil {
			t.Error(err)
		}
		if !result {
			t.Errorf("ShrinkFile(%s, test_data/temp.png) Failed.", datum.inputName)
		}
	}
	for _, datum := range invalidInput {
		err := ShrinkFile(datum.inputName, datum.expectedName)
		if err == nil {
			t.Errorf("ShrinkFile(%s, %s) Should've failed.", datum.inputName, datum.expectedName)
		}
	}
}

func TestShrinkImg(t *testing.T) {
	for _, datum := range testData {
		fsrc, err := os.Open(datum.inputName)
		if err != nil {
			t.Error(err)
		}
		fdst, err := os.Open(datum.expectedName)
		if err != nil {
			t.Error(err)
		}
		input, _, err := image.Decode(fsrc)
		if err != nil {
			t.Error(err)
		}
		expected, _, err := image.Decode(fdst)
		if err != nil {
			t.Error(err)
		}
		result, err := ShrinkImg(input)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ShrinkImg(%s) Failed.", datum.inputName)
		}
	}
	for _, datum := range invalidInput {
		fsrc, err1 := os.Open(datum.inputName)
		fdst, err2 := os.Open(datum.expectedName)
		input, _, err3 := image.Decode(fsrc)
		_, _, err4 := image.Decode(fdst)
		_, err5 := ShrinkImg(input)
		if err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
			t.Errorf("ShrinkFile(%s, %s) Should've failed.", datum.inputName, datum.expectedName)
		}
	}
}
