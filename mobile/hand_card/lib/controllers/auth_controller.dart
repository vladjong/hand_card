import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

import '../pages/sign-up_page.dart';
// import 'package:shared_preferences/shared_preferences.dart';

class AuthController{

  final client = http.Client();

  final Future<SharedPreferences> _prefs = SharedPreferences.getInstance();

  Future signIn(
    String login,
    String password,
    BuildContext context
    ) async {
      // SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
      const url = 'http://10.0.2.2:8080/auth/sign-in';
      final Uri uri = Uri.parse(url);
      final http.Response response = await client.post(uri,
        body: jsonEncode({
          "password":password,
          "login":login
        }),
      );
      final statusCode = response.statusCode;
      final body = response.body;
      if (statusCode==200) {
        final json = jsonDecode(body);
        var token = json['token'];
        print(token);
        final SharedPreferences prefs = await _prefs;
        await prefs.setString('token', token);

        // Navigator.push(context, MaterialPageRoute(builder: (context) => SignUpPage()));
        // Navigator.of(context).pushAndRemoveUntil(
        //   MaterialPageRoute(builder: (BuildContext (context) => SignUpPage()))
        // );
        // Navigator.of(context).pushAndRemo
        print(body);
      } else {
        print("Login Error");
      }
  }
}