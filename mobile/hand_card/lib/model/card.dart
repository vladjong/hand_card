import 'package:flutter/material.dart';

class DiscountCard {
  final String organization;
  final String categoryName;
  final String number;

  const DiscountCard ({
    required this.organization,
    required this.categoryName,
    required this.number,
  });

static DiscountCard fromJson(json) => DiscountCard(
  organization: json['organization'],
  categoryName: json['category_name'],
  number: json['number']);
}
