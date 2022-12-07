import 'package:flutter/material.dart';
import 'package:hand_card/pages/home_page.dart';
import 'package:hand_card/pages/sign-in_page.dart';
import 'package:hand_card/pages/sign-up_page.dart';
import 'package:hand_card/service/user_secure_storage.dart';

Future<void> main() async {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  Future<bool> checkToken() async{
    if (await UserSecureStorage.getJwt() == null) {
      return false;
    }
    return true;
  }
  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: checkToken(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {

        }
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
          '/':(context) => HomePage(),
        },
        initialRoute: '/sign-in',
        // if (snapshot.data && snapshot.hasData) {
        //   initialRoute: '/sign-in',
        // } else {
        //   initialRoute: '/sign-in',
        );
      },
    );
  }
}

