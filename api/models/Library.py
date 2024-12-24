from django.db import models

class Library(models.Model):
    """
    The database model of a Library

    Represent the place where library is and contact infos

    Attributes:
    ----------
    name      : str
        The name of the library
    town      : str
        The town where it is located
    address_1 : str
        The first address line
    address_2 : str
        The second address line
    phone     : str
        The phone number of the library (can be dot separated)
    email     : str
        The email address to contact the library
    about     : str
        A short description of the library and tell about the people and its purpose / atmosphere
    """
    name      = models.CharField(max_length=64)
    town      = models.CharField(max_length=32)
    address_1 = models.CharField(max_length=128)
    address_2 = models.CharField(max_length=128)
    phone     = models.CharField(max_length=16)
    email     = models.CharField(max_length=64)
    about     = models.TextField()

    class Meta:
        ordering = ["name"]
    
    def __str__(self):
        return self.name