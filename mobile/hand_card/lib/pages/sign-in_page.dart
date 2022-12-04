import 'package:flutter/material.dart';
import 'package:hand_card/controllers/auth_controller.dart';
import 'package:hand_card/pages/sign-up_page.dart';

class SignInPage extends StatefulWidget {
  @override
  _SigInPageState createState() => _SigInPageState();
}

class _SigInPageState extends State<SignInPage> {
  AuthController authController = AuthController();
  @override
  Widget build(BuildContext context) {
    final loginController = TextEditingController();
    final passwordController = TextEditingController();
    return Scaffold(
      body: SafeArea(
        child: Center(
          // ignore: prefer_const_literals_to_create_immutables
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
            const Text(
              'HandCard',
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 24
              ),
            ),
            const SizedBox(height: 20),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 25.0),
              child: Container(
              decoration: BoxDecoration(
                color: Colors.grey[200],
                borderRadius: BorderRadius.circular(12),
              ),
              // ignore: prefer_const_constructors
              child: Padding(
                padding: const EdgeInsets.only(left: 20.0),
                child: TextField(
                  controller: loginController,
                  // ignore: prefer_const_constructors
                  decoration: InputDecoration(
                    border: InputBorder.none,
                    hintText: 'Логин',
                  ),
                ),
              ),
              ),
            ),
            const SizedBox(height: 10),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 25.0),
              child: Container(
              decoration: BoxDecoration(
                color: Colors.grey[200],
                borderRadius: BorderRadius.circular(12),
              ),
              child: Padding(
                padding: const EdgeInsets.only(left: 20.0),
                child: TextField(
                  controller: passwordController,
                  obscureText: true,
                  // ignore: prefer_const_constructors
                  decoration: InputDecoration(
                    border: InputBorder.none,
                    hintText: 'Пароль',
                  ),
                ),
              ),
              ),
            ),
            const SizedBox(height: 20),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 25.0),
              child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.deepPurpleAccent,
                  minimumSize: const Size.fromHeight(50),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(12),
                  ),
                ),
                onPressed: (){
                  authController.signIn(loginController.text, passwordController.text, context);
                },
                child: const Text(
                  'Вход',
                  style: TextStyle(fontSize: 20),
                )
              ),
            ),
            const SizedBox(height: 10),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text("Нет аккаунта?"),
                TextButton(
                  onPressed: (){
                    Navigator.push(context, MaterialPageRoute(builder: (context) => SignUpPage()));
                  },
                    child: const Text(
                      'Регистрация',
                      style: TextStyle(
                        color: Colors.deepPurpleAccent,
                      ),
                  ),
                ),
              ],
            ),
          ]),
        ),
      ),
    );
  }
}