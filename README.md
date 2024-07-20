A JSON toolkit for the *arr apps

# Recipes

## Delete movies from on instance that exist in another

```bash
arrson r -u "http://:{API_KEY_1}@{URL_1}" list | jq .tmdbId | xargs -n 1 -o arrson r -u "https//:{API_KEY_2}@{URL_2}" delete
```