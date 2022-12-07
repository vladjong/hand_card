import 'package:flutter/material.dart';
import 'package:hand_card/pages/home_page.dart';
import 'package:hand_card/pages/sign-in_page.dart';
import 'package:hand_card/pages/sign-up_page.dart';
import 'package:hand_card/service/user_secure_storage.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  final auth = await UserSecureStorage.getJwt() == null ? false : true;

  runApp(MyApp(auth: auth));
}

Future<bool> checkToken() async {
  if (await UserSecureStorage.getJwt() == null) {
    return false;
  }
  return true;
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key, required this.auth}) : super(key: key);

  final bool auth;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
          pageTransitionsTheme: const PageTransitionsTheme(builders: {
            TargetPlatform.android: CupertinoPageTransitionsBuilder()
          }),
          primarySwatch: Colors.blue,
          visualDensity: VisualDensity.adaptivePlatformDensity),
      routes: {
        '/sign-in': (context) => SignInPage(),
        '/sign-up': (context) => SignUpPage(),
        '/': (context) => HomePage(),
      },
      initialRoute: auth ? '/' : '/sign-in',
    );
  }
}

