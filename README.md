A toolkit for the arrs

# Recipes

## Delete movies from on instance that exist in another

```bash
arr r -u "http://:{API_KEY_1}@{URL_1}" list | jq .tmdbId | xargs -n 1 arr r -u "https//:{API_KEY_2}@{URL_2}" delete -y
```