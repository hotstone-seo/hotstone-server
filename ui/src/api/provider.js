import client from "./client";

function match(path) {
  return client
    .post("/provider/match-rule", { path })
    .then(response => response.data)
    .catch(error => {
      throw error;
    });
}

function fetchTags(rule, locale, contentData) {
  const { rule_id, path_param } = rule;
  return client
    .post("/provider/tags", {
      rule_id,
      locale,
      path_param,
      data: contentData
    })
    .then(response => response.data)
    .catch(error => {
      throw error;
    });
}

export { match, fetchTags };
