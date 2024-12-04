from django.db import models

# Book related models

class Author(models.Model):
    name    = models.CharField(max_length=32)
    surname = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return "%s %s".format(self.name, self.surname)

class Editor(models.Model):
    name = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name

class BookType(models.Model):
    name = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name

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

# User related models


