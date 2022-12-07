import 'dart:convert';
import 'package:geolocator/geolocator.dart';
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
      backgroundColor: Colors.black87,
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
                        fontSize: 30,
                        color: Colors.redAccent,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    Text(
                      ' карты',
                      style: TextStyle(fontSize: 30,
                      color: Colors.white),
                    ),
                    ],),
                    ElevatedButton(
                      onPressed: (){},
                      style: ElevatedButton.styleFrom(
                        shape: const CircleBorder(),
                        padding: const EdgeInsets.all(7),
                        backgroundColor: Colors.redAccent,
                        foregroundColor: Colors.red,
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
          child: CardWidget(card: card)
      );
    }
  );
}

class CardWidget extends StatelessWidget {
  const CardWidget({super.key, required this.card});

  final DiscountCard card;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
          onTap: () => Scaffold.of(context).
          showBottomSheet((context) => 
            DetailCardWidget(card: card,)
          ),
          child: Card(
            color: Colors.white,
            shape: RoundedRectangleBorder(
              side: const BorderSide(
                color: Colors.redAccent,
                width: 7
              ),
              borderRadius: BorderRadius.circular(12.0),
          ),
          child: Container(
            width: 330,
            height: 200,
            alignment: Alignment.center,
            padding: const EdgeInsets.all(30),
            child: Text(
              card.organization,
              style: const TextStyle(fontSize: 25, fontWeight: FontWeight.w800),
              textAlign: TextAlign.center,
            ),
          ),
          ),
        );
  }
}

class DetailCardWidget extends StatelessWidget {
  const DetailCardWidget({super.key, required this.card});

  final DiscountCard card;
  @override
  Widget build(BuildContext context) {
    return Container(
      height: 470,
       width: 393,
       decoration: const BoxDecoration(
        color: Colors.black87,
        borderRadius: BorderRadius.only(
          topRight: Radius.circular(20),
          topLeft: Radius.circular(20),
          
        )
      ),
      child: Column(
        children: [
          const SizedBox(height: 30,),
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Card(
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(12.0),
              ),
              child: Container(
                width: 350,
                height: 200,
                alignment: Alignment.center,
                padding: const EdgeInsets.all(30),
                child: Text(
                  card.organization,
                  style: const TextStyle(fontSize: 25, fontWeight: FontWeight.w800),
                ),
              ),
            ),
          ),
          const SizedBox(height: 15,),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 25.0),
            child: Container(
              alignment: Alignment.topLeft,
              child: const Text(
                "Данные карты",
                textAlign: TextAlign.left,
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 20,
                  fontWeight: FontWeight.bold,),
              ),
            ),
          ),
          const SizedBox(height: 20,),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 25.0),
            child: Container(
              alignment: Alignment.topLeft,
              child: const Text(
                "Номер карты",
                textAlign: TextAlign.left,
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 16,
                  fontWeight: FontWeight.w300,),
              ),
            ),
          ),
          const SizedBox(height: 15,),
          Text(
            card.number,
            textAlign: TextAlign.left,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 18,
              fontWeight: FontWeight.w500,),
          ),
          const SizedBox(height: 15,),
                    Padding(
            padding: const EdgeInsets.symmetric(horizontal: 25.0),
            child: Container(
              alignment: Alignment.topLeft,
              child: const Text(
                "Категория карты",
                textAlign: TextAlign.left,
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 16,
                  fontWeight: FontWeight.w300,),
              ),
            ),
          ),
          const SizedBox(height: 15,),

          Text(
            card.categoryName,
            textAlign: TextAlign.left,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 18,
              fontWeight: FontWeight.w500,),
          ),
        ]),
    );
  }
}
