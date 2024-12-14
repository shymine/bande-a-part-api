from api.models import Book, Author, Genre, Editor
from datetime import date

author = Author(
    name="Victor",
    surname="Hugo"
)
author.save()

editor = Editor(
    name="Argyll"
)
editor.save()

bookType = Genre(
    name="sci-fi"
)
bookType.save()

print("authors: {}\neditors: {}\nbook types: {}".format(
    Author.objects.all(),
    Editor.objects.all(),
    Genre.objects.all()
))

book = Book(
    name="blue sky",
    parution=date(1997, 8, 31),
    price=38,
    synopsis="blablablablablabla",
    ISBN="1234567891234",
    stock=3,
    new=True,
    note="oui oui c est bien",
    editor=editor
)
book.save()
book.authors.add(author)
book.genres.add(bookType)

print("books: {}".format(
    Book.objects.all()
))

book.save()
print("books: {}".format(
    Book.objects.all()
))