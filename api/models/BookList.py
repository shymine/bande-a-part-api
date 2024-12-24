from django.db import models

from api.models.Book import Book

class BookList(models.Model):
    name        = models.CharField(max_length=64)
    description = models.TextField()
    priority    = models.SmallIntegerField()

    books = models.ManyToManyField(Book)

    class Meta:
        ordering = ["priority"]
    
    def __str__(self):
        return "{} ({})".format(self.name, self.priority)