import 'dart:ffi';

import 'package:flutter/material.dart';

class DiscountCard {
  final int id;
  final String organization;
  final String categoryName;
  final String number;

  const DiscountCard ({
    required this.id,
    required this.organization,
    required this.categoryName,
    required this.number,
  });

static DiscountCard fromJson(json) => DiscountCard(
  id: json['id'],
  organization: json['organization'],
  categoryName: json['category_name'],
  number: json['number']);
}
