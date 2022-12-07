import 'dart:convert';
import 'package:hand_card/model/card.dart';
import 'package:hand_card/pages/sign-in_page.dart';
import 'package:hand_card/service/user_secure_storage.dart';
import 'package:http/http.dart' as http;


import 'package:flutter/material.dart';

import '../controllers/home_controller.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}


class _HomePageState extends State<HomePage> {

  HomeController homeController = HomeController();

  @override
  Widget build(BuildContext context) =>
    Scaffold(
      body: Column(
            children: [
              const SizedBox(height: 30,),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 25.0),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Row(children: const [
                      Text(
                      'Мои',
                      style: TextStyle(
                        fontSize: 28,
                        color: Colors.deepPurpleAccent,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    Text(
                      ' карты',
                      style: TextStyle(fontSize: 28),
                    ),
                    ],),
                    ElevatedButton(
                      onPressed: (){},
                      style: ElevatedButton.styleFrom(
                        shape: const CircleBorder(),
                        padding: const EdgeInsets.all(7),
                        backgroundColor: Colors.deepPurpleAccent,
                        foregroundColor: Colors.purple,
                      ),
                      child: const Icon(Icons.add, color: Colors.white,),
                    ),
                  ]
                ),
              ),
              Expanded(
                child: FutureBuilder<List<DiscountCard>>(
                  future: homeController.getCards(),
                  builder: (context, snapshot) {
                    if (snapshot.hasData) {
                      final cards = snapshot.data!;
                      return buildCards(cards);
                    } else {
                      return const Text("No cards data");
                    }
                  }
              )
            ),
            ]
          ),
      // )
    );

  Widget buildCards(List<DiscountCard> cards) => ListView.builder(
    itemCount: cards.length,
    shrinkWrap: true,
    itemBuilder: (context, index) {
      final card = cards[index];
      return Padding(
        padding: const EdgeInsets.all(20.0),
        child: Card(
          shape: RoundedRectangleBorder(
            side: const BorderSide(
              color: Colors.deepPurpleAccent,
              width: 7
            ),
            borderRadius: BorderRadius.circular(12.0),
        ),
        child: Container(
          width: 330,
          height: 200,
          alignment: Alignment.center,
          padding: const EdgeInsets.all(30),
          child: Text(card.organization,
          style: const TextStyle(fontSize: 25, fontWeight: FontWeight.normal),
          textAlign: TextAlign.center,),
        ),
                // title: Text(card.organization)    // subtitle: Text(card.number),
        ),
      );
    }
  );
}
