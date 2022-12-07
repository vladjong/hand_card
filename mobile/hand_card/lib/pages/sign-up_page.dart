import 'package:flutter/material.dart';

import 'sign-in_page.dart';

class SignUpPage extends StatefulWidget {
  @override
  // ignore: library_private_types_in_public_api
  _SignUpPageState createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  @override
  Widget build(BuildContext context) {
    final emailController = TextEditingController();
    final loginController = TextEditingController();
    final passwordController = TextEditingController();
    final passwordConfirmController = TextEditingController();
    return Scaffold(
      body: SafeArea(
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
            const Text(
              'Регистрация',
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
              child: Padding(
                padding: const EdgeInsets.only(left: 20.0),
                child: TextField(
                  controller: emailController,
                  decoration: const InputDecoration(
                    border: InputBorder.none,
                    hintText: 'Email',
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
              // ignore: prefer_const_constructors
              child: Padding(
                padding: const EdgeInsets.only(left: 20.0),
                child: TextField(
                  controller: loginController,
                  decoration: const InputDecoration(
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
                  // controller: passwordController,
                  obscureText: true,
                  // ignore: prefer_const_constructors
                  controller: passwordController,
                  // ignore: prefer_const_constructors
                  decoration: InputDecoration(
                    border: InputBorder.none,
                    hintText: 'Пароль',
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
                padding: EdgeInsets.only(left: 20.0),
                child: TextField(
                  // controller: passwordController,
                  obscureText: true,
                  // ignore: prefer_const_constructors
                  controller: passwordConfirmController,
                  // ignore: prefer_const_constructors
                  decoration: InputDecoration(
                    border: InputBorder.none,
                    hintText: 'Подтвердить пароль',
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
                  // authController.signIn(loginController.text, passwordController.text);
                },
                child: const Text(
                  'Зарегистрироваться',
                  style: TextStyle(fontSize: 20),
                )
              ),
            ),
            const SizedBox(height: 10),
            TextButton(
              onPressed: (){
                Navigator.of(context).pushNamed('/sign-in');
              },
              child: const Text(
                'Уже есть аккаунт',
                style: TextStyle(
                  color: Colors.deepPurple,
                ),
              ),
            ),
          ]),
        ),
      ),
    );
  }
}