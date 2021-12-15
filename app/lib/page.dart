import 'package:flutter/material.dart';

class SettingsPage extends StatelessWidget {
  const SettingsPage({Key? key, required this.child, required this.title})
      : super(key: key);

  final Widget child;
  final String title;

  @override
  Widget build(BuildContext context) {
    const horizontalDivider = Divider(
      height: 1,
      indent: 0,
      endIndent: 0,
    );

    return Scaffold(
      appBar: AppBar(
        title: const Text('Linux Control Center'),
      ),
      body: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(
            width: 160,
            child: ListView(
              scrollDirection: Axis.vertical,
              shrinkWrap: true,
              children: [
                ListTile(
                  selected: title == 'Lenovo',
                  leading: const Icon(Icons.laptop),
                  title: const Text('Lenovo'),
                  onTap: () {
                    Navigator.of(context).pushReplacementNamed('/lenovo');
                  },
                ),
                horizontalDivider,
              ],
            ),
          ),
          const VerticalDivider(
            width: 1,
            indent: 0,
            endIndent: 0,
          ),
          Expanded(child: child),
        ],
      ),
    );
  }
}
