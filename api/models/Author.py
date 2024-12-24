from django.db import models

class Author(models.Model):
    """
    The database model of an Author

    Attributes:
    ----------
    name    : str
        The name of the author
    surname : str
        The surname of the author
    """
    name    = models.CharField(max_length=32)
    surname = models.CharField(max_length=32)

    class Meta:
        ordering = ["name"]
    
    def __str__(self) -> str:
        return "{} {}".format(self.name, self.surname)