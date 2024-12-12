from django.http import JsonResponse
from django.shortcuts import render
import rest_framework.request

from api.models import Author
from api.serializers.authorSerializers import AuthorSerializer

from django.views.decorators.csrf import csrf_exempt

import rest_framework
# Create your views here.

@csrf_exempt
def list_author(request: rest_framework.request.HttpRequest):
    assert request.method == "GET"
    data = Author.objects.all()
    serializer = AuthorSerializer(data, many=True)
    return JsonResponse(serializer.data, safe=False)

