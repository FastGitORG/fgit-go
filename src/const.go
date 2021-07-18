package main

const (
	mainHelpMsg = "FastGit Command Line Tool\n" +
		"=========================\n" +
		"REMARKS\n" +
		"    We will convert GitHub to FastGit automatically\n" +
		"    Do everything like git\n" +
		"    Build by KevinZonda with GoLang\n" +
		"EXTRA-SYNTAX\n" +
		"    fgit debug [URL<string>] [--help|-h]\n" +
		"    fgit get [URL<string>] [Path<string>] [--help|-h]\n" +
		"    fgit conv [Target<string>] [--help|-h]\n " +
		"    If you want to known more about extra-syntax, try to use --help"
	jsdHelpMsg = "FastGit JsdGet Command Line Tool\n" +
		"=============================\n" +
		"REMARKS\n" +
		"    Download with jsDelivr automatically\n" +
		"SYNTAX\n" +
		"    fgit [--help|-h]\n" +
		"    fgit jsdget [URL<string>]\n" +
		"    fgit jsdget [URL<string>] [Path<string>]\n" +
		"ALIASES\n" +
		"    fgit jdl\n" +
		"EXAMPLE\n" +
		"    fgit jsdget https://github.com/fastgitorg/fgit-go/archive/master.zip"
	getHelpMsg = "FastGit Get Command Line Tool\n" +
		"=============================\n" +
		"REMARKS\n" +
		"    Download with FastGit automatically\n" +
		"SYNTAX\n" +
		"    fgit [--help|-h]\n" +
		"    fgit get [URL<string>]\n" +
		"    fgit get [URL<string>] [Path<string>]\n" +
		"ALIASES\n" +
		"    fgit dl\n" +
		"    fgit download\n" +
		"EXAMPLE\n" +
		"    fgit get https://github.com/fastgitorg/fgit-go/archive/master.zip"
	debugHelpMsg = "FastGit Debug Command Line Tool\n" +
		"===============================\n" +
		"SYNTAX\n" +
		"    fgit debug [URL<string>] [--help|-h]\n" +
		"REMARKS\n" +
		"    URL is an optional parameter\n" +
		"    We debug https://hub.fastgit.org by default\n" +
		"    If you want to debug another URL, enter URL param\n" +
		"EXAMPLE\n" +
		"    fgit debug\n" +
		"    fgit debug https://fastgit.org"
	convHelpMsg = "FastGit Conv Command Line Tool\n" +
		"==============================\n" +
		"REMARKS\n" +
		"    Convert upstream between github or fastgit automatically\n" +
		"    github and gh means convert to github, fastgit and fg means convert to fastgit\n" +
		"SYNTAX\n" +
		"    fgit conv [--help|-h]\n" +
		"    fgit conv [github|gh|fastgit|fg]\n" +
		"ALIASES\n" +
		"    fgit convert\n" +
		"EXAMPLE\n" +
		"    fgit conv gh"
)
