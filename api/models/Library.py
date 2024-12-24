from django.db import models

class Library(models.Model):
    name = models.CharField(max_length=64)
    town = models.CharField(max_length=32)
    address_1 = models.CharField(max_length=128)
    address_2 = models.CharField(max_length=128)
    phone = models.CharField(max_length=16)
    email = models.CharField(max_length=64)
    about = models.TextField()

    class Meta:
        ordering = ["name"]
    
    def __str__(self):
        return self.name