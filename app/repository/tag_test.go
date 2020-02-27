package repository_test

import (
	"testing"

	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"github.com/stretchr/testify/require"
)

func TestTagValidation(t *testing.T) {
	t.Run("When rule ID is missing", func(t *testing.T) {
		tag := repository.Tag{Locale: "en_US", Type: "title", Value: "Some Title"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.RuleID' Error:Field validation for 'RuleID' failed on the 'required' tag")
	})
	t.Run("When locale is missing", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Type: "title", Value: "Some Title"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Locale' Error:Field validation for 'Locale' failed on the 'required' tag")
	})
	t.Run("When type is missing", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Locale: "en_US"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Type' Error:Field validation for 'Type' failed on the 'required' tag")
	})
	t.Run("When valid title tag", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Locale: "en_US", Type: "title", Value: "Some Title"}
		require.NoError(t, tag.Validate())
	})
	t.Run("When value is missing from title tag", func(t *testing.T) {
		tag := repository.Tag{RuleID: 999, Locale: "en_US", Type: "title"}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Value' Error:Field validation for 'Value' failed on the 'noempty' tag")
	})
	t.Run("When valid meta tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "meta",
			Attributes: []byte(`{ "name": "description", "content": "Some description" }`),
		}
		require.NoError(t, tag.Validate())
	})
	t.Run("When 'name' key is missing from attribute in meta tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "meta",
			Attributes: []byte(`{ "content": "Some description" }`),
		}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Attributes' Error:Field validation for 'Attributes' failed on the '' tag")
	})
	t.Run("When 'content' key is missing from attribute in meta tag", func(t *testing.T) {
		tag := repository.Tag{
			RuleID:     999,
			Locale:     "en_US",
			Type:       "meta",
			Attributes: []byte(`{ "name": "description" }`),
		}
		require.EqualError(t, tag.Validate(),
			"Key: 'Tag.Attributes' Error:Field validation for 'Attributes' failed on the '' tag")
	})
}
