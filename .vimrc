call plug#begin()
Plug 'fatih/vim-go', { 'do': ':GoInstallBinaries' }
Plug 'AndrewRadev/splitjoin.vim'
Plug 'SirVer/ultisnips'
Plug 'ctrlpvim/ctrlp.vim'
call plug#end()
autocmd BufNewFile,BufRead *.go setlocal noexpandtab tabstop=4 shiftwidth=4
set autowrite
syntax on
set nocompatible
filetype plugin indent on
set noexpandtab
set tabstop=4
set shiftwidth=4
set wildmode=longest,list
set incsearch
