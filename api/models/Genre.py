from django.db import models

class Genre(models.Model):
    name = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name