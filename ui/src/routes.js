import React from "react";
import MetatagForm from "./views/Metatag/MetatagForm";

const RuleList = React.lazy(() => import("./views/Rule/RuleList"));
const RuleForm = React.lazy(() => import("./views/Rule/RuleForm"));
const RuleDetail = React.lazy(() => import("./views/Rule/RuleDetail"));
const MetatagPreview = React.lazy(() =>
  import("./views/Metatag/MetatagPreview")
);

const DataSourceList = React.lazy(() =>
  import("./views/DataSource/DataSource")
);
const DataSourceForm = React.lazy(() =>
  import("./views/DataSource/DataSourceForm")
);
const CanonicalForm = React.lazy(() =>
  import("./views/Canonical/CanonicalForm")
);
const TitletagForm = React.lazy(() => import("./views/Titletag/TitletagForm"));
const ScripttagForm = React.lazy(() =>
  import("./views/Scripttag/ScripttagForm")
);
const LanguageList = React.lazy(() => import("./views/Language/Language"));
const LanguageForm = React.lazy(() => import("./views/Language/LanguageForm"));

const MismatchRuleList = React.lazy(() =>
  import("./views/MismatchRule/MismatchRuleList")
);

const AnalyticPage = React.lazy(() => import("./views/Analytic/AnalyticPage"));
const SimulationPage = React.lazy(() =>
  import("./views/Simulation/SimulationPage")
);

const routes = [
  { path: "/", exact: true, name: "Home" },
  { path: "/rule", name: "Rules", component: RuleList },
  { path: "/ruleForm", name: "Rules", component: RuleForm },
  {
    path: "/ruleDetail",
    name: "Rule Details",
    component: RuleDetail
  },
  { path: "/metatagForm", name: "Metatag", component: MetatagForm },
  {
    path: "/metatagPreview",
    name: "MetatagPreview",
    component: MetatagPreview
  },
  { path: "/datasource", name: "DataSource", component: DataSourceList },
  { path: "/dataSourceForm", name: "DataSource", component: DataSourceForm },
  { path: "/canonicalForm", name: "Canonical", component: CanonicalForm },
  { path: "/titletagForm", name: "Titletag", component: TitletagForm },
  { path: "/scripttagForm", name: "Scripttag", component: ScripttagForm },
  { path: "/language", name: "Language", component: LanguageList },
  { path: "/languageForm", name: "Language", component: LanguageForm },
  { path: "/mismatchRule", name: "MismatchRule", component: MismatchRuleList },
  { path: "/analytic", name: "Analytic", component: AnalyticPage },
  { path: "/simulation", name: "Simulation", component: SimulationPage }
];

export default routes;
