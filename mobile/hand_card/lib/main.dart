import 'package:flutter/material.dart';
import 'package:hand_card/pages/home_page.dart';
import 'package:hand_card/pages/sign-in_page.dart';
import 'package:hand_card/pages/sign-up_page.dart';


Future<void> main() async {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        pageTransitionsTheme: const PageTransitionsTheme(
          builders: {
            TargetPlatform.android: CupertinoPageTransitionsBuilder()
          }
        ),
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity
      ),
      routes: {
        '/sign-in':(context) => SignInPage(),
        '/sign-up':(context) => SignUpPage(),
      },
      initialRoute: '/sign-in',
      // home: SignInPage(),
    );
  }
}

