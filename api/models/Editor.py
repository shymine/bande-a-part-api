from django.db import models

class Editor(models.Model):
    """
    The database model of an Editor

    Attributes:
    ----------
    name : str
        The name of the editor
    """
    name = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return self.name