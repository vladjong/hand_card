
import 'package:flutter/material.dart';

class SnackBarService {
  static const errorColor = Colors.redAccent;
  static const okColor = Colors.greenAccent;

  static Future<void> showSnackBar(
    BuildContext context, String msg, bool error) async {
      ScaffoldMessenger.of(context).removeCurrentSnackBar();

      final snackBar = SnackBar(
        content: Text(msg),
        backgroundColor: error? errorColor : okColor,
      );
      ScaffoldMessenger.of(context).showSnackBar(snackBar);
    }
}