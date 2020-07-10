@echo off
:: Add-Git-Alias
:: by KevinZonda
:: Thanks @b1f6c1c4 & @zhullyb

:: Install & Import to env git and fgit before using this.

cls

echo ===================================
echo =   Add Git Alias                 =
echo =   by KevinZonda                 =
echo =   Thanks @b1f6c1c4 ^& @zhullyb   =
echo ===================================

echo Install ^& import git and fgit to env before using this.
pause

cls
echo ===================================
echo =   Add Git Alias                 =
echo =   by KevinZonda                 =
echo =   Thanks @b1f6c1c4 ^& @zhullyb   =
echo ===================================

echo -^> Adding...
echo --^> fclone
git config --global alias.fclone '!fgit clone'

echo --^> fpull
git config --global alias.fpull '!fgit pull'

echo --^> debug
git config --global alias.debug '!fgit debug'

echo --^> dl ^& get ^& download
git config --global alias.dl '!fgit dl'
git config --global alias.get '!fgit get'
git config --global alias.download '!fgit download'

echo --^> conv
git config --global alias.conv '!fgit conv'
git config --global alias.convert '!fgit convert'

pause
