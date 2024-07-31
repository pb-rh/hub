{{ comment "handleGRPCServer starts configures and starts a gRPC server on the given URL. It shuts down the server if any error is received in the error channel." }}
func handleGRPCServer(ctx context.Context, u *url.URL{{ range $.Services }}{{ if .Service.Methods }}, {{ .Service.VarName }}Endpoints *{{ .Service.PkgName }}.Endpoints{{ end }}{{ end }}, wg *sync.WaitGroup, errc chan error, dbg bool) {