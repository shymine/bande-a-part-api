from api.models import Library
from api.serializers.LibrarySerializer import LibrarySerializer
from rest_framework import generics

class LibraryListView(generics.ListCreateAPIView):
    queryset = Library.objects.all()
    serializer_class = LibrarySerializer

class LibraryDetail(generics.RetrieveUpdateDestroyAPIView):
    queryset = Library.objects.all()
    serializer_class = LibrarySerializer