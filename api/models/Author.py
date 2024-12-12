from django.db import models


class Author(models.Model):
    name    = models.CharField(max_length=32)
    surname = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return "{} {}".format(self.name, self.surname)