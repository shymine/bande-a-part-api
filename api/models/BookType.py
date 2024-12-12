from django.db import models

class BookType(models.Model):
    name = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name