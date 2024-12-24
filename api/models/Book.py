from django.db import models

from api.models.Author import Author
from api.models.Editor import Editor
from api.models.Genre import Genre

class Book(models.Model):
    """
    The database model for a Book

    Attributes:
    ----------
    name     : str
        The title of the Book
    parution : date
        The date of release of the Book
    price    : positive float
        The price to pay for this book
    synopsis : str
        The synopsis of the book, usually found on the book backpage
    ISBN     : 13 digits int
        The Identification number of the book
    stock    : positive int
        The number of remaining books in stock
    new      : boolean
        Is this book considered new (if yes, will be put up front)
    note     : str
        The note of the librarian about this book
    authors  : []Author
        The authors of the book
    editor   : Editor
        The editor of the book
    genres   : []Genre
        The genres of the book (mainly for tagging) such as sci-fi, horror, etc.
    """
    name     = models.CharField(max_length=64)
    parution = models.DateField()
    price    = models.FloatField()
    synopsis = models.TextField()
    ISBN     = models.CharField(max_length=13)
    stock    = models.IntegerField()
    new      = models.BooleanField()
    note     = models.TextField()

    authors  = models.ManyToManyField(Author)
    editor   = models.ForeignKey(Editor, on_delete=models.DO_NOTHING)
    genres    = models.ManyToManyField(Genre)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name