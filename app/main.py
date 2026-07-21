import streamlit as st
import time
from rec_book_client import BookClient, Book
from functions import cadastro, lista_simples, THEMES, GENRES, update_book_prices
from amazon_client import AmazonClient

import pandas as pd

st.set_page_config(page_title="Recomendação de Livros", page_icon=":books:", layout="wide")

st.markdown("# Recomendação de Livros")

book_client = BookClient(base_url="http://localhost:3030")
amazon_client = AmazonClient()

cadastro_tab, listas_simples_tab = st.tabs(["Cadastro de Livros", "Listas Simples"])

with cadastro_tab:
    cadastro(book_client)
    
    if st.button("Atualizar preços dos livros"):
        with st.spinner("Atualizando preços..."):
            update_book_prices(book_client, amazon_client)


with listas_simples_tab:
    lista_simples(book_client)
    
