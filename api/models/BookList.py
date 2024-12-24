from django.db import models

from api.models.Book import Book

class BookList(models.Model):
    """
    The database model of a BookList

    This list can be for promotionnal events, for themed package of books, etc.

    Attributes:
    ----------
    name        : str
        The name of the book list
    description : str
        A description of this list
    priority    : positive int
        The priority of this list against the other lists, determine the order they will be displayed
    books       : []Book
        The books composing the list
   """
    name        = models.CharField(max_length=64)
    description = models.TextField()
    priority    = models.SmallIntegerField()

    books = models.ManyToManyField(Book)

    class Meta:
        ordering = ["priority"]
    
    def __str__(self):
        return "{} ({})".format(self.name, self.priority)