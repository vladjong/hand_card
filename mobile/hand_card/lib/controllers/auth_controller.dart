import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:hand_card/pages/home_page.dart';
import 'package:hand_card/service/snack_bar.dart';
import 'package:http/http.dart' as http;

import '../pages/sign-up_page.dart';
import '../service/user_secure_storage.dart';

class AuthController{

  final client = http.Client();

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
         // ignore: use_build_context_synchronously
        await UserSecureStorage.setJwt(token);
        Navigator.push(context, MaterialPageRoute(builder: (context) => HomePage()));
      } else {
        // ignore: use_build_context_synchronously
        SnackBarService.showSnackBar(
          context,
          'Неправильный логин или пароль. Повторите попытку',
          true,
        );
      }
  }
}