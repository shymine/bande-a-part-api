from django.db import models

class Genre(models.Model):
    """
    The database model of a Genre

    Attributes:
    ----------
    name : str
        The name of the Genre (sci-fi, horror, etc.)
    """
    name = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name