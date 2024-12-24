from api.models import BookList
from api.serializers.bookListSerializer import BookListSerializer
from rest_framework import generics

class BookListListView(generics.ListCreateAPIView):
    queryset = BookList.objects.all()
    serializer_class = BookListSerializer

class BookListDetail(generics.RetrieveUpdateDestroyAPIView):
    queryset = BookList.objects.all()
    serializer_class = BookListSerializer
