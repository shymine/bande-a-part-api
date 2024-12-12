from django.db import models

from api.models.Author import Author
from api.models.Editor import Editor
from api.models.BookType import BookType

class Book(models.Model):
    name     = models.CharField(max_length=64)
    parution = models.DateField()
    price    = models.FloatField()
    synopsis = models.TextField()
    ISBN     = models.CharField(max_length=13)
    stock    = models.IntegerField()
    new      = models.BooleanField()
    note     = models.TextField()

    author   = models.ManyToManyField(Author)
    editor   = models.ForeignKey(Editor, on_delete=models.DO_NOTHING)
    types    = models.ManyToManyField(BookType)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name