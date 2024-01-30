import 'package:http/http.dart' as http;
import 'dart:io';
import 'dart:async';
import 'dart:convert';

class ApiProvider {
  ApiProvider();

  String endPoint = 'http://192.168.21.90:8080';

  Future<http.Response> doLogin(String username, String password) async {
    var url = '$endPoint/login';

    var body = {
      "username": username,
      "password": password,
    };

    // var connect = await http.post(Uri.parse(_Uri), body: body);
    var response = await http.post(Uri.parse(url),
        headers: {
          HttpHeaders.contentTypeHeader: 'application/json',
        },
        body: jsonEncode(body)
    );
    return response;
  }
}
