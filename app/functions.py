# %%

import streamlit as st

import pandas as pd

from rec_book_client import Book, BookClient
from amazon_client import AmazonClient

import time

# %%


## Generos
GENRES = [
    "Dados",
    "Tecnologia",
    "Gestão",
    "Negócios",
]

GENRES.sort()

# Temas
THEMES = [
    "Data Science",
    "Data Engineering",
    "Data Analytics",
    "Machine Learning",
    "Deep Learning",
    "Business Intelligence",
    "Gestão de Projetos",
    "Gestão de Pessoas",
    "Gestão de Negócios",
    "Programação",
    "Desenvolvimento Web",
    "Desenvolvimento Mobile",
    "Arquitetura de Software",
    "Engenharia de Software",
    "Banco de Dados",
    "Segurança da Informação",
    "Inteligência Artificial",
    "Computação em Nuvem",
    "DevOps",
    "Marketing",
    "CRM",
    "Finanças",
    "Logística",
    "Análise de Crédito",
    "Gestão de Riscos",
    "Anti-Fraude",
]

THEMES.sort()


def cadastro(book_client):

    books = book_client.get_books()
    books = [Book(id=None, title="Cadastrar Novo Livro", author="", description="", genre="", theme="", link="")] + books

    book = st.selectbox(
        "Selecione um livro para obter recomendações:",
        options=books,
        format_func=lambda book: f"{book.title} - {book.author}"
    )

    if book.title != "Cadastrar Novo Livro":
        edit_toggle = st.toggle("Editar Livro", key="edit_book_button")
    else:
        edit_toggle = True


    title_c, author_c = st.columns(2)

    title = title_c.text_input("Título", value=book.title, disabled=not edit_toggle)
    author = author_c.text_input("Autor", value=book.author, disabled=not edit_toggle)

    description = st.text_area("Descrição", value=book.description, disabled=not edit_toggle)

    genre_c, theme_c, link_c, price_c, price_discount_c = st.columns(5)
    genre = genre_c.selectbox("Gênero", options=GENRES, index=GENRES.index(book.genre) if book.genre in GENRES else 0, disabled=not edit_toggle)
    theme = theme_c.selectbox("Tema", options=THEMES, index=THEMES.index(book.theme) if book.theme in THEMES else 0, disabled=not edit_toggle)
    link = link_c.text_input("Link", value=book.link, disabled=not edit_toggle)
    price = price_c.number_input("Preço", value=book.price if book.price is not None else 0.01, disabled=not edit_toggle, min_value=0.00, format="%.2f")
    price_with_discount = price_discount_c.number_input("Preço com Desconto", value=book.price_with_discount if book.price_with_discount is not None else 0.01, disabled=not edit_toggle, min_value=0.00, format="%.2f")


    if edit_toggle:
        save_button = st.button("Salvar Alterações", disabled=not edit_toggle)
        if save_button:
            
            if book.id is None:
                new_book = Book(id=None,
                                title=title,
                                author=author,
                                description=description,
                                genre=genre,
                                theme=theme,
                                link=link,
                                price=price,
                                price_with_discount=price_with_discount)
                res=book_client.create_book(new_book)
                
                if "error" in res:
                    st.error(res["error"])
                
                else:
                    st.success("Livro cadastrado com sucesso!")
                    time.sleep(1)
                    st.rerun()
            
            else:
                updated_book = Book(id=book.id,title=title,author=author,description=description,genre=genre,theme=theme,link=link, price=price, price_with_discount=price_with_discount)
                res=book_client.update_book(updated_book)
                
                if "error" in res:
                    st.error(res["error"])
                
                else:
                    st.success("Livro atualizado com sucesso!")
                    time.sleep(1)
                    st.rerun()


def lista_simples(book_client):

    st.write("Listas Simples")
    
    genre_c, theme_c = st.columns(2)
    genre = genre_c.selectbox("Selecione um gênero:", options=GENRES+["Todos"], index=len(GENRES), key="genre_select")
    theme = theme_c.selectbox("Selecione um tema:", options=THEMES+["Todos"], index=len(THEMES), key="theme_select")
    
    params = {}
    if theme != "Todos":
        params["theme"] = theme
    if genre != "Todos":
        params["genre"] = genre
        
    books = book_client.search_books(params)
    
    books = [b.to_dict() for b in books]  # Convert Book objects to dictionaries for display
    df_book = pd.DataFrame(books)
    
    df_book["discount_percentage"] = 1 - (df_book["price_with_discount"] / df_book["price"])
    df_book["discount_percentage"] = df_book["discount_percentage"].fillna(0)
    
    st.write(f"Total de livros encontrados: {len(books)}")
    
    if len(books) == 0:
        st.warning("Nenhum livro encontrado com os filtros selecionados.")
    
    else:
        df_columns = {
            "title": st.column_config.TextColumn("Título", help="Título do livro", width="large"),
            "author": st.column_config.TextColumn("Autor", help="Autor do livro", width="medium"),
            "description": st.column_config.TextColumn("Descrição", help="Descrição do livro", width="large"),
            "genre": st.column_config.TextColumn("Gênero", help="Gênero do livro", width="medium"),
            "theme": st.column_config.TextColumn("Tema", help="Tema do livro", width="medium"),
            "link": st.column_config.LinkColumn("Link", help="Clique para acessar o livro", width="medium"),
            "price": st.column_config.NumberColumn("Preço", help="Preço do livro", width="small", format="%.2f"),
            "price_with_discount": st.column_config.NumberColumn("Preço com Desconto", help="Preço com desconto do livro", width="small", format="%.2f"),
            "discount_percentage": st.column_config.NumberColumn("Desconto", help="Percentual de desconto", width="small", format="percent"),
        }
        
        df_book = df_book.drop("id", axis=1)
        
        st.dataframe(df_book, hide_index=True, use_container_width=True, column_config=df_columns)
        
        if st.toggle("Exportar Markdown"):
            
            for g in sorted(df_book["genre"].unique()):
                df_genre = df_book[df_book["genre"] == g]
                st.markdown(f"### {g}")
                
                for t in sorted(df_genre["theme"].unique()):
                    df_theme = df_genre[df_genre["theme"] == t]
                    st.markdown(f"#### {t}")
                    
                    for _, row in df_theme.iterrows():
                        st.markdown(f"- **{row['title']}**: [{row['link']}]({row['link']})")
                        

# %%

def get_book_info(book: Book, amazon_client):
    
    sort_by = [
        None,
        "Relevance",
        "AvgCustomerReviews",
        "Featured",
        "NewestArrivals",
        "Price:HighToLow",
        "Price:LowToHigh",
    ]

    for s in sort_by:
        result = amazon_client.search_items(book=book, search_index="Books", sort_by=s)
        
        try:
            item = result.get("searchResult",{}).get("items", [])[0]
            item_dimension = (item.get("itemInfo", {})
                                  .get("productInfo", {})
                                  .get("itemDimensions", None))
            
            if item_dimension is None:
                continue
            
            price = item.get("offersV2", {}).get("listings", [])[0].get("price", None)
            
            price_with_discount = price.get('money', {}).get('amount') if price else None
            full_price = price.get('savingBasis', {}).get('money', {}).get('amount', price_with_discount) if price else None
            
            if price_with_discount is None or full_price is None:
                continue
            
            book.price_with_discount = price_with_discount
            book.price = full_price
            book.link = item.get("detailPageURL", book.link) 
            
            return book

        except IndexError as e:
            print(f"Error processing result for sort_by={s}: {e}")
            continue


def update_book_prices(book_client, amazon_client):
    
    books = book_client.get_books()
    
    bar_progress= st.progress(0, text="Atualizando preços dos livros...")
    
    total_books = len(books)
    
    for i, book in enumerate(books):
        updated_book = get_book_info(book=book, amazon_client=amazon_client)
        
        if updated_book:
            res = book_client.update_book(updated_book)
            
            if "error" in res:
                st.error(f"Erro ao atualizar o livro {book.title}: {res['error']}")
            else:
                st.success(f"Livro {book.title} atualizado com sucesso!")
        else:
            st.warning(f"Não foi possível obter o preço do livro {book.title}.")
        
        bar_progress.progress((i + 1) / total_books, text=f"Atualizando preços dos livros... ({i + 1}/{total_books})")
        time.sleep(1)
