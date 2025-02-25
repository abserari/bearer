package integration_test

import "testing"

const javascriptRulesPath string = "../../pkg/commands/process/settings/rules/javascript/"

func TestJavascriptLangLogger(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/logger")
}

func TestJavascriptOpenRedirect(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/open_redirect")
}

func TestJavascriptLangSession(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/session")
}

func TestJavascriptWeakEncryption(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/weak_encryption")
}

func TestJavascriptWeakPasswordEncryption(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/weak_password_encryption")
}

func TestJavascriptJWT(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/jwt")
}

func TestJavascriptJWTWeakEncryption(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/jwt_weak_encryption")
}

func TestJavascriptJWTHardcodedSecret(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/jwt_hardcoded_secret")
}

func TestJavascriptHTTPInsecure(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/http_insecure")
}

func TestJavascriptLangException(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/exception")
}

func TestJavascriptLangFileGeneration(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/file_generation")
}

func TestJavascriptHardcodedSecret(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/hardcoded_secret")
}

func TestJavascriptDangeoursInsertHTML(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"lang/dangerous_insert_html")
}

func TestJavascriptAwsLambdaCodeInjection(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"aws_lambda/code_injection")
}

func TestJavascriptAwsLambdaQueryInjection(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"aws_lambda/query_injection")
}

func TestJavascriptAwsLambdaSqlInjection(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"aws_lambda/sql_injection")
}

func TestJavascriptAwsLambdaOsCommandInjection(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"aws_lambda/os_command_injection")
}

func TestJavascriptExpressHardCodedSecret(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/hardcoded_secret")
}

func TestJavascriptExpressOpenRedirect(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/open_redirect")
}

func TestJavascriptExpressUnsafeDeserialization(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/unsafe_deserialization")
}

func TestJavascriptExpressExternalResource(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/external_resource")
}

func TestJavascriptExpressInsecureAllowOrigin(t *testing.T) {
	getRunner(t).runTest(t, javascriptRulesPath+"express/insecure_allow_origin")
}

func TestJavascriptExpressExternalFileUpload(t *testing.T) {
	getRunner(t).runTest(t, javascriptRulesPath+"express/external_file_upload")
}

func TestJavascriptExpressJwtNotRevoked(t *testing.T) {
	getRunner(t).runTest(t, javascriptRulesPath+"express/jwt_not_revoked")
}

func TestJavascriptExpressExposedDirListing(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/exposed_dir_listing")
}

func TestJavascriptExpressCrossSiteScripting(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/cross_site_scripting")
}

func TestJavascriptExpressPathTraversal(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/path_traversal")
}

func TestJavascriptExpressHttpsProtocolMissing(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/https_protocol_missing")
}

func TestJavascriptExpressServerSideRequestForgery(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/server_side_request_forgery")
}

func TestJavascriptExpressInsecureTemplateRendering(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/insecure_template_rendering")
}

func TestJavascriptExpressStaticAssetWithSession(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/static_asset_with_session")
}

func TestJavascriptExpressUiRedress(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/ui_redress")
}

func TestJavascriptExpressSqlInjection(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/sql_injection")
}

func TestJavascriptExpressSecureCookie(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/insecure_cookie")
}

func TestJavascriptExpressDefaultCookieConfig(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/default_cookie_config")
}

func TestJavascriptExpressDefaultSessionConfig(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/default_session_config")
}

func TestJavascriptExpressXXEVulnerability(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/xml_external_entity_vulnerability")
}

func TestJavascriptExpressEvalUserInput(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/eval_user_input")
}

func TestJavascriptReactGoogleAnalytics(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"react/google_analytics")
}

func TestJavascriptReactDangerouslySetInnerHTML(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"react/dangerously_set_inner_html")
}

func TestJavascriptThirdPartySentry(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/sentry")
}

func TestJavascriptThirdPartyOpenAI(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/openai")
}

func TestJavascriptGTM(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/google_tag_manager")
}

func TestJavascriptGoogleAnalytics(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/google_analytics")
}

func TestJavascriptAlgolia(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/algolia")
}

func TestJavascriptDataDog(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/datadog")
}

func TestJavascriptDataDogBrowser(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/datadog_browser")
}

func TestJavascriptElasticSearch(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/elasticsearch")
}

func TestJavascriptSegmentDataflow(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/segment")
}

func TestJavascriptNewRelic(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/new_relic")
}

func TestJavascriptRollbar(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/rollbar")
}

func TestJavascriptHoneybadger(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/honeybadger")
}

func TestJavascriptAirbrake(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/airbrake")
}

func TestJavascriptOpenTelemetry(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/open_telemetry")
}

func TestJavascriptBugsnag(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/bugsnag")
}

func TestJavascripPassportHardcodedSecret(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/passport_hardcoded_secret")
}

func TestJavascriptDomPurify(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"third_parties/dom_purify")
}

func TestJavascriptHelmetMissing(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/helmet_missing")
}

func TestJavascriptReduceFingerprint(t *testing.T) {
	t.Parallel()
	getRunner(t).runTest(t, javascriptRulesPath+"express/reduce_fingerprint")
}
